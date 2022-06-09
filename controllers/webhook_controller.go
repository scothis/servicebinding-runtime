/*
Copyright 2022 Scott Andrews.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/go-logr/logr"
	"github.com/vmware-labs/reconciler-runtime/reconcilers"
	"github.com/vmware-labs/reconciler-runtime/tracker"
	"gomodules.xyz/jsonpatch/v2"
	admissionv1 "k8s.io/api/admission/v1"
	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/util/workqueue"
	controllerruntime "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	servicebindingv1beta1 "github.com/scothis/servicebinding-runtime/apis/v1beta1"
	"github.com/scothis/servicebinding-runtime/projector"
	"github.com/scothis/servicebinding-runtime/rbac"
	"github.com/scothis/servicebinding-runtime/resolver"
)

//+kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=mutatingwebhookconfigurations,verbs=get;list;watch;create;update;patch
//+kubebuilder:rbac:groups=core,resources=events,verbs=get;list;watch;create;update;patch;delete

// AdmissionProjector reconciles a MutatingWebhookConfiguration object
func AdmissionProjectorReconciler(c reconcilers.Config, name string, accessChecker rbac.AccessChecker) *reconcilers.AggregateReconciler {
	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name: name,
		},
	}

	return &reconcilers.AggregateReconciler{
		Name:    "AdmissionProjector",
		Type:    &admissionregistrationv1.MutatingWebhookConfiguration{},
		Request: req,
		Reconciler: reconcilers.Sequence{
			LoadServiceBindings(req),
			InterceptGVKs(),
			WebhookRules([]admissionregistrationv1.OperationType{admissionregistrationv1.Create, admissionregistrationv1.Update}, accessChecker),
		},
		DesiredResource: func(ctx context.Context, resource *admissionregistrationv1.MutatingWebhookConfiguration) (client.Object, error) {
			if resource == nil || len(resource.Webhooks) != 1 {
				// the webhook config isn't in a form that we expect, ignore it
				return resource, nil
			}
			rules := RetrieveWebhookRules(ctx)
			resource.Webhooks[0].Rules = rules
			return resource, nil
		},
		SemanticEquals: func(a1, a2 *admissionregistrationv1.MutatingWebhookConfiguration) bool {
			if a1 == nil || len(a1.Webhooks) != 1 || a2 == nil || len(a2.Webhooks) != 1 {
				// the webhook config isn't in a form that we expect, ignore it
				return true
			}
			return equality.Semantic.DeepEqual(a1.Webhooks[0].Rules, a2.Webhooks[0].Rules)
		},
		MergeBeforeUpdate: func(current, desired *admissionregistrationv1.MutatingWebhookConfiguration) {
			if current == nil || len(current.Webhooks) != 1 || desired == nil || len(desired.Webhooks) != 1 {
				// the webhook config isn't in a form that we expect, ignore it
				return
			}
			current.Webhooks[0].Rules = desired.Webhooks[0].Rules
		},
		Sanitize: func(resource *admissionregistrationv1.MutatingWebhookConfiguration) []admissionregistrationv1.RuleWithOperations {
			if resource == nil || len(resource.Webhooks) == 0 {
				return nil
			}
			return resource.Webhooks[0].Rules
		},

		Setup: func(ctx context.Context, mgr controllerruntime.Manager, bldr *builder.Builder) error {
			if err := mgr.GetFieldIndexer().IndexField(ctx, &servicebindingv1beta1.ServiceBinding{}, workloadRefIndexKey, func(obj client.Object) []string {
				serviceBinding := obj.(*servicebindingv1beta1.ServiceBinding)
				gvk := schema.FromAPIVersionAndKind(serviceBinding.Spec.Workload.APIVersion, serviceBinding.Spec.Workload.Kind)
				return []string{workloadRefIndexValue(gvk.Group, gvk.Kind)}
			}); err != nil {
				return err
			}
			return nil
		},
		Config: c,
	}
}

func AdmissionProjectorWebhook(c reconcilers.Config) *admission.Webhook {
	return &admission.Webhook{
		Handler: admission.HandlerFunc(func(ctx context.Context, r admission.Request) admission.Response {
			response := admission.Response{
				AdmissionResponse: admissionv1.AdmissionResponse{
					UID: r.UID,
				},
			}

			workload := &unstructured.Unstructured{}
			if err := json.Unmarshal(r.Object.Raw, workload); err != nil {
				// TODO handle error
				return response
			}

			// find matching service bindings
			serviceBindings := &servicebindingv1beta1.ServiceBindingList{}
			if err := c.List(ctx, serviceBindings, client.InNamespace(r.Namespace), client.MatchingFields{workloadRefIndexKey: workloadRefIndexValue(r.Kind.Group, r.Kind.Kind)}); err != nil {
				// TODO handle error
				return response
			}

			// check that bindings are for this workload
			activeServiceBindings := []servicebindingv1beta1.ServiceBinding{}
			for _, sb := range serviceBindings.Items {
				if !sb.DeletionTimestamp.IsZero() {
					continue
				}
				w := sb.Spec.Workload
				if w.Name == r.Name {
					activeServiceBindings = append(activeServiceBindings, sb)
					continue
				}
				if w.Selector != nil {
					selector, err := metav1.LabelSelectorAsSelector(w.Selector)
					if err != nil {
						// TODO handle error
						return response
					}
					if selector.Matches(labels.Set(workload.GetLabels())) {
						activeServiceBindings = append(activeServiceBindings, sb)
						continue
					}
				}
			}

			// project active bindings into workload
			projector := projector.New(resolver.New(c))
			projectedWorkload := workload.DeepCopy()
			for i := range activeServiceBindings {
				sb := activeServiceBindings[i].DeepCopy()
				sb.Default()
				if err := projector.Project(ctx, sb, projectedWorkload); err != nil {
					// TODO handle error
					return response
				}
			}

			if !equality.Semantic.DeepEqual(workload, projectedWorkload) {
				// add patch to response

				workloadBytes, err := json.Marshal(workload)
				if err != nil {
					// TODO handle error
					return response
				}
				projectedWorkloadBytes, err := json.Marshal(projectedWorkload)
				if err != nil {
					// TODO handle error
					return response
				}
				patch, err := jsonpatch.CreatePatch(workloadBytes, projectedWorkloadBytes)
				if err != nil {
					// TODO handle error
					return response
				}
				response.Patches = patch
			}

			response.Allowed = true
			return response
		}),
	}
}

//+kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations,verbs=get;list;watch;create;update;patch
//+kubebuilder:rbac:groups=core,resources=events,verbs=get;list;watch;create;update;patch;delete

// TriggerReconciler reconciles a ValidatingWebhookConfiguration object
func TriggerReconciler(c reconcilers.Config, name string, accessChecker rbac.AccessChecker) *reconcilers.AggregateReconciler {
	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name: name,
		},
	}

	return &reconcilers.AggregateReconciler{
		Name:    "Trigger",
		Type:    &admissionregistrationv1.ValidatingWebhookConfiguration{},
		Request: req,
		Reconciler: reconcilers.Sequence{
			LoadServiceBindings(req),
			TriggerGVKs(),
			InterceptGVKs(),
			WebhookRules([]admissionregistrationv1.OperationType{admissionregistrationv1.Create, admissionregistrationv1.Update, admissionregistrationv1.Delete}, accessChecker),
		},
		DesiredResource: func(ctx context.Context, resource *admissionregistrationv1.ValidatingWebhookConfiguration) (client.Object, error) {
			if resource == nil || len(resource.Webhooks) != 1 {
				// the webhook config isn't in a form that we expect, ignore it
				return resource, nil
			}
			rules := RetrieveWebhookRules(ctx)
			resource.Webhooks[0].Rules = rules
			return resource, nil
		},
		SemanticEquals: func(a1, a2 *admissionregistrationv1.ValidatingWebhookConfiguration) bool {
			if a1 == nil || len(a1.Webhooks) != 1 || a2 == nil || len(a2.Webhooks) != 1 {
				// the webhook config isn't in a form that we expect, ignore it
				return true
			}
			return equality.Semantic.DeepEqual(a1.Webhooks[0].Rules, a2.Webhooks[0].Rules)
		},
		MergeBeforeUpdate: func(current, desired *admissionregistrationv1.ValidatingWebhookConfiguration) {
			if current == nil || len(current.Webhooks) != 1 || desired == nil || len(desired.Webhooks) != 1 {
				// the webhook config isn't in a form that we expect, ignore it
				return
			}
			current.Webhooks[0].Rules = desired.Webhooks[0].Rules
		},
		Sanitize: func(resource *admissionregistrationv1.ValidatingWebhookConfiguration) []admissionregistrationv1.RuleWithOperations {
			if resource == nil || len(resource.Webhooks) == 0 {
				return nil
			}
			return resource.Webhooks[0].Rules
		},

		Config: c,
	}
}

func TriggerWebhook(c reconcilers.Config, serviceBindingController controller.Controller) *admission.Webhook {
	return &admission.Webhook{
		Handler: admission.HandlerFunc(func(ctx context.Context, r admission.Request) admission.Response {
			log := logr.FromContextOrDiscard(ctx)

			response := admission.Response{
				AdmissionResponse: admissionv1.AdmissionResponse{
					UID:     r.UID,
					Allowed: true,
				},
			}

			// TODO find a better way to get at the queue, this is fragile and may break in any controller-runtime update
			queue := reflect.ValueOf(serviceBindingController).Elem().FieldByName("Queue").Interface().(workqueue.Interface)

			trackKey := tracker.NewKey(
				schema.GroupVersionKind{
					Group:   r.Kind.Group,
					Version: r.Kind.Version,
					Kind:    r.Kind.Kind,
				},
				types.NamespacedName{
					Namespace: r.Namespace,
					Name:      r.Name,
				},
			)
			for _, nsn := range c.Tracker.Lookup(ctx, trackKey) {
				req := reconcile.Request{NamespacedName: nsn}
				log.V(2).Info("enqueue tracked request", "request", req, "for", trackKey, "dryRun", r.DryRun)
				if !(r.DryRun != nil && *r.DryRun) {
					// ignore dry run requests
					queue.Add(req)
				}
			}

			return response
		}),
	}
}

func LoadServiceBindings(req reconcile.Request) reconcilers.SubReconciler {
	return &reconcilers.SyncReconciler{
		Name: "LoadServiceBindings",
		Sync: func(ctx context.Context, _ client.Object) error {
			c := reconcilers.RetrieveConfigOrDie(ctx)

			serviceBindings := &servicebindingv1beta1.ServiceBindingList{}
			if err := c.List(ctx, serviceBindings); err != nil {
				return err
			}

			StashServiceBindings(ctx, serviceBindings.Items)

			return nil
		},
		Setup: func(ctx context.Context, mgr controllerruntime.Manager, bldr *builder.Builder) error {
			bldr.Watches(&source.Kind{Type: &servicebindingv1beta1.ServiceBinding{}}, handler.EnqueueRequestsFromMapFunc(
				func(o client.Object) []reconcile.Request {
					return []reconcile.Request{req}
				},
			))
			return nil
		},
	}
}

func InterceptGVKs() reconcilers.SubReconciler {
	return &reconcilers.SyncReconciler{
		Name: "InterceptGVKs",
		Sync: func(ctx context.Context, _ client.Object) error {
			serviceBindings := RetrieveServiceBindings(ctx)
			gvks := RetrieveObservedGKVs(ctx)

			for i := range serviceBindings {
				workload := serviceBindings[i].Spec.Workload
				gvk := schema.FromAPIVersionAndKind(workload.APIVersion, workload.Kind)
				gvks = append(gvks, gvk)
			}

			StashObservedGVKs(ctx, gvks)

			return nil
		},
	}
}

func TriggerGVKs() reconcilers.SubReconciler {
	return &reconcilers.SyncReconciler{
		Name: "TriggerGVKs",
		Sync: func(ctx context.Context, _ client.Object) error {
			serviceBindings := RetrieveServiceBindings(ctx)
			gvks := RetrieveObservedGKVs(ctx)

			for i := range serviceBindings {
				service := serviceBindings[i].Spec.Service
				gvk := schema.FromAPIVersionAndKind(service.APIVersion, service.Kind)
				if gvk.Kind == "Secret" && (gvk.Group == "" || gvk.Group == "core") {
					// ignore direct bindings
					continue
				}
				gvks = append(gvks, gvk)
			}

			StashObservedGVKs(ctx, gvks)

			return nil
		},
	}
}

func WebhookRules(operations []admissionregistrationv1.OperationType, accessChecker rbac.AccessChecker) reconcilers.SubReconciler {
	return &reconcilers.SyncReconciler{
		Name: "WebhookRules",
		Sync: func(ctx context.Context, _ client.Object) error {
			log := logr.FromContextOrDiscard(ctx)
			c := reconcilers.RetrieveConfigOrDie(ctx)

			// dedup gvks as gvrs
			gvks := RetrieveObservedGKVs(ctx)
			groupResources := map[string]map[string]interface{}{}
			for _, gvk := range gvks {
				rm, err := c.RESTMapper().RESTMapping(gvk.GroupKind(), gvk.Version)
				if err != nil {
					return err
				}
				gvr := rm.Resource
				if _, ok := groupResources[gvr.Group]; !ok {
					groupResources[gvr.Group] = map[string]interface{}{}
				}
				groupResources[gvr.Group][gvr.Resource] = true
			}

			// normalize rules to a canonical form
			rules := []admissionregistrationv1.RuleWithOperations{}
			groups := sets.NewString()
			for group := range groupResources {
				groups.Insert(group)
			}
			for _, group := range groups.List() {
				resources := sets.NewString()
				for resource := range groupResources[group] {
					resources.Insert(resource)
				}

				// check that we have permission to interact with these resources. Admission webhooks bypass RBAC
				for _, resource := range resources.List() {
					if !accessChecker.CanI(ctx, group, resource) {
						log.Info("ignoring resource, access denied", "group", group, "resource", resource)
						resources.Delete(resource)
					}
				}

				if resources.Len() == 0 {
					continue
				}

				rules = append(rules, admissionregistrationv1.RuleWithOperations{
					Operations: operations,
					Rule: admissionregistrationv1.Rule{
						APIGroups:   []string{group},
						APIVersions: []string{"*"},
						Resources:   resources.List(),
					},
				})
			}

			StashWebhookRules(ctx, rules)

			return nil
		},
	}
}

const ServiceBindingsStashKey reconcilers.StashKey = "servicebinding.io:servicebindings"

func StashServiceBindings(ctx context.Context, serviceBindings []servicebindingv1beta1.ServiceBinding) {
	reconcilers.StashValue(ctx, ServiceBindingsStashKey, serviceBindings)
}

func RetrieveServiceBindings(ctx context.Context) []servicebindingv1beta1.ServiceBinding {
	value := reconcilers.RetrieveValue(ctx, ServiceBindingsStashKey)
	if serviceBindings, ok := value.([]servicebindingv1beta1.ServiceBinding); ok {
		return serviceBindings
	}
	return nil
}

const ObservedGVKsStashKey reconcilers.StashKey = "servicebinding.io:observedgvks"

func StashObservedGVKs(ctx context.Context, gvks []schema.GroupVersionKind) {
	reconcilers.StashValue(ctx, ObservedGVKsStashKey, gvks)
}

func RetrieveObservedGKVs(ctx context.Context) []schema.GroupVersionKind {
	value := reconcilers.RetrieveValue(ctx, ObservedGVKsStashKey)
	if refs, ok := value.([]schema.GroupVersionKind); ok {
		return refs
	}
	return nil
}

const WebhookRulesStashKey reconcilers.StashKey = "servicebinding.io:webhookrules"

func StashWebhookRules(ctx context.Context, rules []admissionregistrationv1.RuleWithOperations) {
	reconcilers.StashValue(ctx, WebhookRulesStashKey, rules)
}

func RetrieveWebhookRules(ctx context.Context) []admissionregistrationv1.RuleWithOperations {
	value := reconcilers.RetrieveValue(ctx, WebhookRulesStashKey)
	if rules, ok := value.([]admissionregistrationv1.RuleWithOperations); ok {
		return rules
	}
	return nil
}

const workloadRefIndexKey = ".metadata.workloadRef"

func workloadRefIndexValue(group, kind string) string {
	return schema.GroupKind{Group: group, Kind: kind}.String()
}
