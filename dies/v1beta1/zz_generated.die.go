//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by diegen. DO NOT EDIT.

package v1beta1

import (
	"dies.dev/apis/meta/v1"
	json "encoding/json"
	fmtx "fmt"
	apiv1beta1 "github.com/scothis/servicebinding-runtime/api/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

var ClusterWorkloadResourceMappingBlank = (&ClusterWorkloadResourceMappingDie{}).DieFeed(apiv1beta1.ClusterWorkloadResourceMapping{})

type ClusterWorkloadResourceMappingDie struct {
	v1.FrozenObjectMeta
	mutable bool
	r       apiv1beta1.ClusterWorkloadResourceMapping
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *ClusterWorkloadResourceMappingDie) DieImmutable(immutable bool) *ClusterWorkloadResourceMappingDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *ClusterWorkloadResourceMappingDie) DieFeed(r apiv1beta1.ClusterWorkloadResourceMapping) *ClusterWorkloadResourceMappingDie {
	if d.mutable {
		d.FrozenObjectMeta = v1.FreezeObjectMeta(r.ObjectMeta)
		d.r = r
		return d
	}
	return &ClusterWorkloadResourceMappingDie{
		FrozenObjectMeta: v1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *ClusterWorkloadResourceMappingDie) DieFeedPtr(r *apiv1beta1.ClusterWorkloadResourceMapping) *ClusterWorkloadResourceMappingDie {
	if r == nil {
		r = &apiv1beta1.ClusterWorkloadResourceMapping{}
	}
	return d.DieFeed(*r)
}

// DieRelease returns the resource managed by the die.
func (d *ClusterWorkloadResourceMappingDie) DieRelease() apiv1beta1.ClusterWorkloadResourceMapping {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *ClusterWorkloadResourceMappingDie) DieReleasePtr() *apiv1beta1.ClusterWorkloadResourceMapping {
	r := d.DieRelease()
	return &r
}

// DieReleaseUnstructured returns the resource managed by the die as an unstructured object.
func (d *ClusterWorkloadResourceMappingDie) DieReleaseUnstructured() runtime.Unstructured {
	r := d.DieReleasePtr()
	u, _ := runtime.DefaultUnstructuredConverter.ToUnstructured(r)
	return &unstructured.Unstructured{
		Object: u,
	}
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *ClusterWorkloadResourceMappingDie) DieStamp(fn func(r *apiv1beta1.ClusterWorkloadResourceMapping)) *ClusterWorkloadResourceMappingDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *ClusterWorkloadResourceMappingDie) DeepCopy() *ClusterWorkloadResourceMappingDie {
	r := *d.r.DeepCopy()
	return &ClusterWorkloadResourceMappingDie{
		FrozenObjectMeta: v1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

var _ runtime.Object = (*ClusterWorkloadResourceMappingDie)(nil)

func (d *ClusterWorkloadResourceMappingDie) DeepCopyObject() runtime.Object {
	return d.r.DeepCopy()
}

func (d *ClusterWorkloadResourceMappingDie) GetObjectKind() schema.ObjectKind {
	r := d.DieRelease()
	return r.GetObjectKind()
}

func (d *ClusterWorkloadResourceMappingDie) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.r)
}

func (d *ClusterWorkloadResourceMappingDie) UnmarshalJSON(b []byte) error {
	if d == ClusterWorkloadResourceMappingBlank {
		return fmtx.Errorf("cannot unmarshal into the blank die, create a copy first")
	}
	if !d.mutable {
		return fmtx.Errorf("cannot unmarshal into immutable dies, create a mutable version first")
	}
	r := &apiv1beta1.ClusterWorkloadResourceMapping{}
	err := json.Unmarshal(b, r)
	*d = *d.DieFeed(*r)
	return err
}

// MetadataDie stamps the resource's ObjectMeta field with a mutable die.
func (d *ClusterWorkloadResourceMappingDie) MetadataDie(fn func(d *v1.ObjectMetaDie)) *ClusterWorkloadResourceMappingDie {
	return d.DieStamp(func(r *apiv1beta1.ClusterWorkloadResourceMapping) {
		d := v1.ObjectMetaBlank.DieImmutable(false).DieFeed(r.ObjectMeta)
		fn(d)
		r.ObjectMeta = d.DieRelease()
	})
}

// SpecDie stamps the resource's spec field with a mutable die.
func (d *ClusterWorkloadResourceMappingDie) SpecDie(fn func(d *ClusterWorkloadResourceMappingSpecDie)) *ClusterWorkloadResourceMappingDie {
	return d.DieStamp(func(r *apiv1beta1.ClusterWorkloadResourceMapping) {
		d := ClusterWorkloadResourceMappingSpecBlank.DieImmutable(false).DieFeed(r.Spec)
		fn(d)
		r.Spec = d.DieRelease()
	})
}

func (d *ClusterWorkloadResourceMappingDie) Spec(v apiv1beta1.ClusterWorkloadResourceMappingSpec) *ClusterWorkloadResourceMappingDie {
	return d.DieStamp(func(r *apiv1beta1.ClusterWorkloadResourceMapping) {
		r.Spec = v
	})
}

var ClusterWorkloadResourceMappingSpecBlank = (&ClusterWorkloadResourceMappingSpecDie{}).DieFeed(apiv1beta1.ClusterWorkloadResourceMappingSpec{})

type ClusterWorkloadResourceMappingSpecDie struct {
	mutable bool
	r       apiv1beta1.ClusterWorkloadResourceMappingSpec
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *ClusterWorkloadResourceMappingSpecDie) DieImmutable(immutable bool) *ClusterWorkloadResourceMappingSpecDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *ClusterWorkloadResourceMappingSpecDie) DieFeed(r apiv1beta1.ClusterWorkloadResourceMappingSpec) *ClusterWorkloadResourceMappingSpecDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &ClusterWorkloadResourceMappingSpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *ClusterWorkloadResourceMappingSpecDie) DieFeedPtr(r *apiv1beta1.ClusterWorkloadResourceMappingSpec) *ClusterWorkloadResourceMappingSpecDie {
	if r == nil {
		r = &apiv1beta1.ClusterWorkloadResourceMappingSpec{}
	}
	return d.DieFeed(*r)
}

// DieRelease returns the resource managed by the die.
func (d *ClusterWorkloadResourceMappingSpecDie) DieRelease() apiv1beta1.ClusterWorkloadResourceMappingSpec {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *ClusterWorkloadResourceMappingSpecDie) DieReleasePtr() *apiv1beta1.ClusterWorkloadResourceMappingSpec {
	r := d.DieRelease()
	return &r
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *ClusterWorkloadResourceMappingSpecDie) DieStamp(fn func(r *apiv1beta1.ClusterWorkloadResourceMappingSpec)) *ClusterWorkloadResourceMappingSpecDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *ClusterWorkloadResourceMappingSpecDie) DeepCopy() *ClusterWorkloadResourceMappingSpecDie {
	r := *d.r.DeepCopy()
	return &ClusterWorkloadResourceMappingSpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

// Versions is the collection of versions for a given resource, with mappings.
func (d *ClusterWorkloadResourceMappingSpecDie) Versions(v ...apiv1beta1.ClusterWorkloadResourceMappingTemplate) *ClusterWorkloadResourceMappingSpecDie {
	return d.DieStamp(func(r *apiv1beta1.ClusterWorkloadResourceMappingSpec) {
		r.Versions = v
	})
}

var ClusterWorkloadResourceMappingTemplateBlank = (&ClusterWorkloadResourceMappingTemplateDie{}).DieFeed(apiv1beta1.ClusterWorkloadResourceMappingTemplate{})

type ClusterWorkloadResourceMappingTemplateDie struct {
	mutable bool
	r       apiv1beta1.ClusterWorkloadResourceMappingTemplate
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *ClusterWorkloadResourceMappingTemplateDie) DieImmutable(immutable bool) *ClusterWorkloadResourceMappingTemplateDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *ClusterWorkloadResourceMappingTemplateDie) DieFeed(r apiv1beta1.ClusterWorkloadResourceMappingTemplate) *ClusterWorkloadResourceMappingTemplateDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &ClusterWorkloadResourceMappingTemplateDie{
		mutable: d.mutable,
		r:       r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *ClusterWorkloadResourceMappingTemplateDie) DieFeedPtr(r *apiv1beta1.ClusterWorkloadResourceMappingTemplate) *ClusterWorkloadResourceMappingTemplateDie {
	if r == nil {
		r = &apiv1beta1.ClusterWorkloadResourceMappingTemplate{}
	}
	return d.DieFeed(*r)
}

// DieRelease returns the resource managed by the die.
func (d *ClusterWorkloadResourceMappingTemplateDie) DieRelease() apiv1beta1.ClusterWorkloadResourceMappingTemplate {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *ClusterWorkloadResourceMappingTemplateDie) DieReleasePtr() *apiv1beta1.ClusterWorkloadResourceMappingTemplate {
	r := d.DieRelease()
	return &r
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *ClusterWorkloadResourceMappingTemplateDie) DieStamp(fn func(r *apiv1beta1.ClusterWorkloadResourceMappingTemplate)) *ClusterWorkloadResourceMappingTemplateDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *ClusterWorkloadResourceMappingTemplateDie) DeepCopy() *ClusterWorkloadResourceMappingTemplateDie {
	r := *d.r.DeepCopy()
	return &ClusterWorkloadResourceMappingTemplateDie{
		mutable: d.mutable,
		r:       r,
	}
}

// Version is the version of the workload resource that this mapping is for.
func (d *ClusterWorkloadResourceMappingTemplateDie) Version(v string) *ClusterWorkloadResourceMappingTemplateDie {
	return d.DieStamp(func(r *apiv1beta1.ClusterWorkloadResourceMappingTemplate) {
		r.Version = v
	})
}

// Annotations is a Restricted JSONPath that references the annotations map within the workload resource. These annotations must end up in the resulting Pod, and are generally not the workload resource's annotations. Defaults to `.spec.template.metadata.annotations`.
func (d *ClusterWorkloadResourceMappingTemplateDie) Annotations(v string) *ClusterWorkloadResourceMappingTemplateDie {
	return d.DieStamp(func(r *apiv1beta1.ClusterWorkloadResourceMappingTemplate) {
		r.Annotations = v
	})
}

// Containers is the collection of mappings to container-like fragments of the workload resource. Defaults to mappings appropriate for a PodSpecable resource.
func (d *ClusterWorkloadResourceMappingTemplateDie) Containers(v ...apiv1beta1.ClusterWorkloadResourceMappingContainer) *ClusterWorkloadResourceMappingTemplateDie {
	return d.DieStamp(func(r *apiv1beta1.ClusterWorkloadResourceMappingTemplate) {
		r.Containers = v
	})
}

// Volumes is a Restricted JSONPath that references the slice of volumes within the workload resource. Defaults to `.spec.template.spec.volumes`.
func (d *ClusterWorkloadResourceMappingTemplateDie) Volumes(v string) *ClusterWorkloadResourceMappingTemplateDie {
	return d.DieStamp(func(r *apiv1beta1.ClusterWorkloadResourceMappingTemplate) {
		r.Volumes = v
	})
}

var ClusterWorkloadResourceMappingContainerBlank = (&ClusterWorkloadResourceMappingContainerDie{}).DieFeed(apiv1beta1.ClusterWorkloadResourceMappingContainer{})

type ClusterWorkloadResourceMappingContainerDie struct {
	mutable bool
	r       apiv1beta1.ClusterWorkloadResourceMappingContainer
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *ClusterWorkloadResourceMappingContainerDie) DieImmutable(immutable bool) *ClusterWorkloadResourceMappingContainerDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *ClusterWorkloadResourceMappingContainerDie) DieFeed(r apiv1beta1.ClusterWorkloadResourceMappingContainer) *ClusterWorkloadResourceMappingContainerDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &ClusterWorkloadResourceMappingContainerDie{
		mutable: d.mutable,
		r:       r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *ClusterWorkloadResourceMappingContainerDie) DieFeedPtr(r *apiv1beta1.ClusterWorkloadResourceMappingContainer) *ClusterWorkloadResourceMappingContainerDie {
	if r == nil {
		r = &apiv1beta1.ClusterWorkloadResourceMappingContainer{}
	}
	return d.DieFeed(*r)
}

// DieRelease returns the resource managed by the die.
func (d *ClusterWorkloadResourceMappingContainerDie) DieRelease() apiv1beta1.ClusterWorkloadResourceMappingContainer {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *ClusterWorkloadResourceMappingContainerDie) DieReleasePtr() *apiv1beta1.ClusterWorkloadResourceMappingContainer {
	r := d.DieRelease()
	return &r
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *ClusterWorkloadResourceMappingContainerDie) DieStamp(fn func(r *apiv1beta1.ClusterWorkloadResourceMappingContainer)) *ClusterWorkloadResourceMappingContainerDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *ClusterWorkloadResourceMappingContainerDie) DeepCopy() *ClusterWorkloadResourceMappingContainerDie {
	r := *d.r.DeepCopy()
	return &ClusterWorkloadResourceMappingContainerDie{
		mutable: d.mutable,
		r:       r,
	}
}

// Path is the JSONPath within the workload resource that matches an existing fragment that is container-like.
func (d *ClusterWorkloadResourceMappingContainerDie) Path(v string) *ClusterWorkloadResourceMappingContainerDie {
	return d.DieStamp(func(r *apiv1beta1.ClusterWorkloadResourceMappingContainer) {
		r.Path = v
	})
}

// Name is a Restricted JSONPath that references the name of the container with the container-like workload resource fragment. If not defined, container name filtering is ignored.
func (d *ClusterWorkloadResourceMappingContainerDie) Name(v string) *ClusterWorkloadResourceMappingContainerDie {
	return d.DieStamp(func(r *apiv1beta1.ClusterWorkloadResourceMappingContainer) {
		r.Name = v
	})
}

// Env is a Restricted JSONPath that references the slice of environment variables for the container with the container-like workload resource fragment. The referenced location is created if it does not exist. Defaults to `.envs`.
func (d *ClusterWorkloadResourceMappingContainerDie) Env(v string) *ClusterWorkloadResourceMappingContainerDie {
	return d.DieStamp(func(r *apiv1beta1.ClusterWorkloadResourceMappingContainer) {
		r.Env = v
	})
}

// VolumeMounts is a Restricted JSONPath that references the slice of volume mounts for the container with the container-like workload resource fragment. The referenced location is created if it does not exist. Defaults to `.volumeMounts`.
func (d *ClusterWorkloadResourceMappingContainerDie) VolumeMounts(v string) *ClusterWorkloadResourceMappingContainerDie {
	return d.DieStamp(func(r *apiv1beta1.ClusterWorkloadResourceMappingContainer) {
		r.VolumeMounts = v
	})
}

var ServiceBindingBlank = (&ServiceBindingDie{}).DieFeed(apiv1beta1.ServiceBinding{})

type ServiceBindingDie struct {
	v1.FrozenObjectMeta
	mutable bool
	r       apiv1beta1.ServiceBinding
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *ServiceBindingDie) DieImmutable(immutable bool) *ServiceBindingDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *ServiceBindingDie) DieFeed(r apiv1beta1.ServiceBinding) *ServiceBindingDie {
	if d.mutable {
		d.FrozenObjectMeta = v1.FreezeObjectMeta(r.ObjectMeta)
		d.r = r
		return d
	}
	return &ServiceBindingDie{
		FrozenObjectMeta: v1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *ServiceBindingDie) DieFeedPtr(r *apiv1beta1.ServiceBinding) *ServiceBindingDie {
	if r == nil {
		r = &apiv1beta1.ServiceBinding{}
	}
	return d.DieFeed(*r)
}

// DieRelease returns the resource managed by the die.
func (d *ServiceBindingDie) DieRelease() apiv1beta1.ServiceBinding {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *ServiceBindingDie) DieReleasePtr() *apiv1beta1.ServiceBinding {
	r := d.DieRelease()
	return &r
}

// DieReleaseUnstructured returns the resource managed by the die as an unstructured object.
func (d *ServiceBindingDie) DieReleaseUnstructured() runtime.Unstructured {
	r := d.DieReleasePtr()
	u, _ := runtime.DefaultUnstructuredConverter.ToUnstructured(r)
	return &unstructured.Unstructured{
		Object: u,
	}
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *ServiceBindingDie) DieStamp(fn func(r *apiv1beta1.ServiceBinding)) *ServiceBindingDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *ServiceBindingDie) DeepCopy() *ServiceBindingDie {
	r := *d.r.DeepCopy()
	return &ServiceBindingDie{
		FrozenObjectMeta: v1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

var _ runtime.Object = (*ServiceBindingDie)(nil)

func (d *ServiceBindingDie) DeepCopyObject() runtime.Object {
	return d.r.DeepCopy()
}

func (d *ServiceBindingDie) GetObjectKind() schema.ObjectKind {
	r := d.DieRelease()
	return r.GetObjectKind()
}

func (d *ServiceBindingDie) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.r)
}

func (d *ServiceBindingDie) UnmarshalJSON(b []byte) error {
	if d == ServiceBindingBlank {
		return fmtx.Errorf("cannot unmarshal into the blank die, create a copy first")
	}
	if !d.mutable {
		return fmtx.Errorf("cannot unmarshal into immutable dies, create a mutable version first")
	}
	r := &apiv1beta1.ServiceBinding{}
	err := json.Unmarshal(b, r)
	*d = *d.DieFeed(*r)
	return err
}

// MetadataDie stamps the resource's ObjectMeta field with a mutable die.
func (d *ServiceBindingDie) MetadataDie(fn func(d *v1.ObjectMetaDie)) *ServiceBindingDie {
	return d.DieStamp(func(r *apiv1beta1.ServiceBinding) {
		d := v1.ObjectMetaBlank.DieImmutable(false).DieFeed(r.ObjectMeta)
		fn(d)
		r.ObjectMeta = d.DieRelease()
	})
}

// SpecDie stamps the resource's spec field with a mutable die.
func (d *ServiceBindingDie) SpecDie(fn func(d *ServiceBindingSpecDie)) *ServiceBindingDie {
	return d.DieStamp(func(r *apiv1beta1.ServiceBinding) {
		d := ServiceBindingSpecBlank.DieImmutable(false).DieFeed(r.Spec)
		fn(d)
		r.Spec = d.DieRelease()
	})
}

// StatusDie stamps the resource's status field with a mutable die.
func (d *ServiceBindingDie) StatusDie(fn func(d *ServiceBindingStatusDie)) *ServiceBindingDie {
	return d.DieStamp(func(r *apiv1beta1.ServiceBinding) {
		d := ServiceBindingStatusBlank.DieImmutable(false).DieFeed(r.Status)
		fn(d)
		r.Status = d.DieRelease()
	})
}

func (d *ServiceBindingDie) Spec(v apiv1beta1.ServiceBindingSpec) *ServiceBindingDie {
	return d.DieStamp(func(r *apiv1beta1.ServiceBinding) {
		r.Spec = v
	})
}

func (d *ServiceBindingDie) Status(v apiv1beta1.ServiceBindingStatus) *ServiceBindingDie {
	return d.DieStamp(func(r *apiv1beta1.ServiceBinding) {
		r.Status = v
	})
}

var ServiceBindingSpecBlank = (&ServiceBindingSpecDie{}).DieFeed(apiv1beta1.ServiceBindingSpec{})

type ServiceBindingSpecDie struct {
	mutable bool
	r       apiv1beta1.ServiceBindingSpec
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *ServiceBindingSpecDie) DieImmutable(immutable bool) *ServiceBindingSpecDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *ServiceBindingSpecDie) DieFeed(r apiv1beta1.ServiceBindingSpec) *ServiceBindingSpecDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &ServiceBindingSpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *ServiceBindingSpecDie) DieFeedPtr(r *apiv1beta1.ServiceBindingSpec) *ServiceBindingSpecDie {
	if r == nil {
		r = &apiv1beta1.ServiceBindingSpec{}
	}
	return d.DieFeed(*r)
}

// DieRelease returns the resource managed by the die.
func (d *ServiceBindingSpecDie) DieRelease() apiv1beta1.ServiceBindingSpec {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *ServiceBindingSpecDie) DieReleasePtr() *apiv1beta1.ServiceBindingSpec {
	r := d.DieRelease()
	return &r
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *ServiceBindingSpecDie) DieStamp(fn func(r *apiv1beta1.ServiceBindingSpec)) *ServiceBindingSpecDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *ServiceBindingSpecDie) DeepCopy() *ServiceBindingSpecDie {
	r := *d.r.DeepCopy()
	return &ServiceBindingSpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

// Name is the name of the service as projected into the workload container.  Defaults to .metadata.name.
func (d *ServiceBindingSpecDie) Name(v string) *ServiceBindingSpecDie {
	return d.DieStamp(func(r *apiv1beta1.ServiceBindingSpec) {
		r.Name = v
	})
}

// Type is the type of the service as projected into the workload container
func (d *ServiceBindingSpecDie) Type(v string) *ServiceBindingSpecDie {
	return d.DieStamp(func(r *apiv1beta1.ServiceBindingSpec) {
		r.Type = v
	})
}

// Provider is the provider of the service as projected into the workload container
func (d *ServiceBindingSpecDie) Provider(v string) *ServiceBindingSpecDie {
	return d.DieStamp(func(r *apiv1beta1.ServiceBindingSpec) {
		r.Provider = v
	})
}

// Workload is a reference to an object
func (d *ServiceBindingSpecDie) Workload(v apiv1beta1.ServiceBindingWorkloadReference) *ServiceBindingSpecDie {
	return d.DieStamp(func(r *apiv1beta1.ServiceBindingSpec) {
		r.Workload = v
	})
}

// Service is a reference to an object that fulfills the ProvisionedService duck type
func (d *ServiceBindingSpecDie) Service(v apiv1beta1.ServiceBindingServiceReference) *ServiceBindingSpecDie {
	return d.DieStamp(func(r *apiv1beta1.ServiceBindingSpec) {
		r.Service = v
	})
}

// Env is the collection of mappings from Secret entries to environment variables
func (d *ServiceBindingSpecDie) Env(v ...apiv1beta1.EnvMapping) *ServiceBindingSpecDie {
	return d.DieStamp(func(r *apiv1beta1.ServiceBindingSpec) {
		r.Env = v
	})
}

var ServiceBindingWorkloadReferenceBlank = (&ServiceBindingWorkloadReferenceDie{}).DieFeed(apiv1beta1.ServiceBindingWorkloadReference{})

type ServiceBindingWorkloadReferenceDie struct {
	mutable bool
	r       apiv1beta1.ServiceBindingWorkloadReference
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *ServiceBindingWorkloadReferenceDie) DieImmutable(immutable bool) *ServiceBindingWorkloadReferenceDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *ServiceBindingWorkloadReferenceDie) DieFeed(r apiv1beta1.ServiceBindingWorkloadReference) *ServiceBindingWorkloadReferenceDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &ServiceBindingWorkloadReferenceDie{
		mutable: d.mutable,
		r:       r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *ServiceBindingWorkloadReferenceDie) DieFeedPtr(r *apiv1beta1.ServiceBindingWorkloadReference) *ServiceBindingWorkloadReferenceDie {
	if r == nil {
		r = &apiv1beta1.ServiceBindingWorkloadReference{}
	}
	return d.DieFeed(*r)
}

// DieRelease returns the resource managed by the die.
func (d *ServiceBindingWorkloadReferenceDie) DieRelease() apiv1beta1.ServiceBindingWorkloadReference {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *ServiceBindingWorkloadReferenceDie) DieReleasePtr() *apiv1beta1.ServiceBindingWorkloadReference {
	r := d.DieRelease()
	return &r
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *ServiceBindingWorkloadReferenceDie) DieStamp(fn func(r *apiv1beta1.ServiceBindingWorkloadReference)) *ServiceBindingWorkloadReferenceDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *ServiceBindingWorkloadReferenceDie) DeepCopy() *ServiceBindingWorkloadReferenceDie {
	r := *d.r.DeepCopy()
	return &ServiceBindingWorkloadReferenceDie{
		mutable: d.mutable,
		r:       r,
	}
}

// API version of the referent.
func (d *ServiceBindingWorkloadReferenceDie) APIVersion(v string) *ServiceBindingWorkloadReferenceDie {
	return d.DieStamp(func(r *apiv1beta1.ServiceBindingWorkloadReference) {
		r.APIVersion = v
	})
}

// Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (d *ServiceBindingWorkloadReferenceDie) Kind(v string) *ServiceBindingWorkloadReferenceDie {
	return d.DieStamp(func(r *apiv1beta1.ServiceBindingWorkloadReference) {
		r.Kind = v
	})
}

// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
func (d *ServiceBindingWorkloadReferenceDie) Name(v string) *ServiceBindingWorkloadReferenceDie {
	return d.DieStamp(func(r *apiv1beta1.ServiceBindingWorkloadReference) {
		r.Name = v
	})
}

// Selector is a query that selects the workload or workloads to bind the service to
func (d *ServiceBindingWorkloadReferenceDie) Selector(v *metav1.LabelSelector) *ServiceBindingWorkloadReferenceDie {
	return d.DieStamp(func(r *apiv1beta1.ServiceBindingWorkloadReference) {
		r.Selector = v
	})
}

// Containers describes which containers in a Pod should be bound to
func (d *ServiceBindingWorkloadReferenceDie) Containers(v ...string) *ServiceBindingWorkloadReferenceDie {
	return d.DieStamp(func(r *apiv1beta1.ServiceBindingWorkloadReference) {
		r.Containers = v
	})
}

var ServiceBindingServiceReferenceBlank = (&ServiceBindingServiceReferenceDie{}).DieFeed(apiv1beta1.ServiceBindingServiceReference{})

type ServiceBindingServiceReferenceDie struct {
	mutable bool
	r       apiv1beta1.ServiceBindingServiceReference
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *ServiceBindingServiceReferenceDie) DieImmutable(immutable bool) *ServiceBindingServiceReferenceDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *ServiceBindingServiceReferenceDie) DieFeed(r apiv1beta1.ServiceBindingServiceReference) *ServiceBindingServiceReferenceDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &ServiceBindingServiceReferenceDie{
		mutable: d.mutable,
		r:       r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *ServiceBindingServiceReferenceDie) DieFeedPtr(r *apiv1beta1.ServiceBindingServiceReference) *ServiceBindingServiceReferenceDie {
	if r == nil {
		r = &apiv1beta1.ServiceBindingServiceReference{}
	}
	return d.DieFeed(*r)
}

// DieRelease returns the resource managed by the die.
func (d *ServiceBindingServiceReferenceDie) DieRelease() apiv1beta1.ServiceBindingServiceReference {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *ServiceBindingServiceReferenceDie) DieReleasePtr() *apiv1beta1.ServiceBindingServiceReference {
	r := d.DieRelease()
	return &r
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *ServiceBindingServiceReferenceDie) DieStamp(fn func(r *apiv1beta1.ServiceBindingServiceReference)) *ServiceBindingServiceReferenceDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *ServiceBindingServiceReferenceDie) DeepCopy() *ServiceBindingServiceReferenceDie {
	r := *d.r.DeepCopy()
	return &ServiceBindingServiceReferenceDie{
		mutable: d.mutable,
		r:       r,
	}
}

// API version of the referent.
func (d *ServiceBindingServiceReferenceDie) APIVersion(v string) *ServiceBindingServiceReferenceDie {
	return d.DieStamp(func(r *apiv1beta1.ServiceBindingServiceReference) {
		r.APIVersion = v
	})
}

// Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (d *ServiceBindingServiceReferenceDie) Kind(v string) *ServiceBindingServiceReferenceDie {
	return d.DieStamp(func(r *apiv1beta1.ServiceBindingServiceReference) {
		r.Kind = v
	})
}

// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
func (d *ServiceBindingServiceReferenceDie) Name(v string) *ServiceBindingServiceReferenceDie {
	return d.DieStamp(func(r *apiv1beta1.ServiceBindingServiceReference) {
		r.Name = v
	})
}

var EnvMappingBlank = (&EnvMappingDie{}).DieFeed(apiv1beta1.EnvMapping{})

type EnvMappingDie struct {
	mutable bool
	r       apiv1beta1.EnvMapping
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *EnvMappingDie) DieImmutable(immutable bool) *EnvMappingDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *EnvMappingDie) DieFeed(r apiv1beta1.EnvMapping) *EnvMappingDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &EnvMappingDie{
		mutable: d.mutable,
		r:       r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *EnvMappingDie) DieFeedPtr(r *apiv1beta1.EnvMapping) *EnvMappingDie {
	if r == nil {
		r = &apiv1beta1.EnvMapping{}
	}
	return d.DieFeed(*r)
}

// DieRelease returns the resource managed by the die.
func (d *EnvMappingDie) DieRelease() apiv1beta1.EnvMapping {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *EnvMappingDie) DieReleasePtr() *apiv1beta1.EnvMapping {
	r := d.DieRelease()
	return &r
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *EnvMappingDie) DieStamp(fn func(r *apiv1beta1.EnvMapping)) *EnvMappingDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *EnvMappingDie) DeepCopy() *EnvMappingDie {
	r := *d.r.DeepCopy()
	return &EnvMappingDie{
		mutable: d.mutable,
		r:       r,
	}
}

// Name is the name of the environment variable
func (d *EnvMappingDie) Name(v string) *EnvMappingDie {
	return d.DieStamp(func(r *apiv1beta1.EnvMapping) {
		r.Name = v
	})
}

// Key is the key in the Secret that will be exposed
func (d *EnvMappingDie) Key(v string) *EnvMappingDie {
	return d.DieStamp(func(r *apiv1beta1.EnvMapping) {
		r.Key = v
	})
}

var ServiceBindingStatusBlank = (&ServiceBindingStatusDie{}).DieFeed(apiv1beta1.ServiceBindingStatus{})

type ServiceBindingStatusDie struct {
	mutable bool
	r       apiv1beta1.ServiceBindingStatus
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *ServiceBindingStatusDie) DieImmutable(immutable bool) *ServiceBindingStatusDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *ServiceBindingStatusDie) DieFeed(r apiv1beta1.ServiceBindingStatus) *ServiceBindingStatusDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &ServiceBindingStatusDie{
		mutable: d.mutable,
		r:       r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *ServiceBindingStatusDie) DieFeedPtr(r *apiv1beta1.ServiceBindingStatus) *ServiceBindingStatusDie {
	if r == nil {
		r = &apiv1beta1.ServiceBindingStatus{}
	}
	return d.DieFeed(*r)
}

// DieRelease returns the resource managed by the die.
func (d *ServiceBindingStatusDie) DieRelease() apiv1beta1.ServiceBindingStatus {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *ServiceBindingStatusDie) DieReleasePtr() *apiv1beta1.ServiceBindingStatus {
	r := d.DieRelease()
	return &r
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *ServiceBindingStatusDie) DieStamp(fn func(r *apiv1beta1.ServiceBindingStatus)) *ServiceBindingStatusDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *ServiceBindingStatusDie) DeepCopy() *ServiceBindingStatusDie {
	r := *d.r.DeepCopy()
	return &ServiceBindingStatusDie{
		mutable: d.mutable,
		r:       r,
	}
}

// ObservedGeneration is the 'Generation' of the ServiceBinding that was last processed by the controller.
func (d *ServiceBindingStatusDie) ObservedGeneration(v int64) *ServiceBindingStatusDie {
	return d.DieStamp(func(r *apiv1beta1.ServiceBindingStatus) {
		r.ObservedGeneration = v
	})
}

// Conditions are the conditions of this ServiceBinding
func (d *ServiceBindingStatusDie) Conditions(v ...metav1.Condition) *ServiceBindingStatusDie {
	return d.DieStamp(func(r *apiv1beta1.ServiceBindingStatus) {
		r.Conditions = v
	})
}

// Binding exposes the projected secret for this ServiceBinding
func (d *ServiceBindingStatusDie) Binding(v *apiv1beta1.ServiceBindingSecretReference) *ServiceBindingStatusDie {
	return d.DieStamp(func(r *apiv1beta1.ServiceBindingStatus) {
		r.Binding = v
	})
}

var ServiceBindingSecretReferenceBlank = (&ServiceBindingSecretReferenceDie{}).DieFeed(apiv1beta1.ServiceBindingSecretReference{})

type ServiceBindingSecretReferenceDie struct {
	mutable bool
	r       apiv1beta1.ServiceBindingSecretReference
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *ServiceBindingSecretReferenceDie) DieImmutable(immutable bool) *ServiceBindingSecretReferenceDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *ServiceBindingSecretReferenceDie) DieFeed(r apiv1beta1.ServiceBindingSecretReference) *ServiceBindingSecretReferenceDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &ServiceBindingSecretReferenceDie{
		mutable: d.mutable,
		r:       r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *ServiceBindingSecretReferenceDie) DieFeedPtr(r *apiv1beta1.ServiceBindingSecretReference) *ServiceBindingSecretReferenceDie {
	if r == nil {
		r = &apiv1beta1.ServiceBindingSecretReference{}
	}
	return d.DieFeed(*r)
}

// DieRelease returns the resource managed by the die.
func (d *ServiceBindingSecretReferenceDie) DieRelease() apiv1beta1.ServiceBindingSecretReference {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *ServiceBindingSecretReferenceDie) DieReleasePtr() *apiv1beta1.ServiceBindingSecretReference {
	r := d.DieRelease()
	return &r
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *ServiceBindingSecretReferenceDie) DieStamp(fn func(r *apiv1beta1.ServiceBindingSecretReference)) *ServiceBindingSecretReferenceDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *ServiceBindingSecretReferenceDie) DeepCopy() *ServiceBindingSecretReferenceDie {
	r := *d.r.DeepCopy()
	return &ServiceBindingSecretReferenceDie{
		mutable: d.mutable,
		r:       r,
	}
}

// Name of the referent secret. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
func (d *ServiceBindingSecretReferenceDie) Name(v string) *ServiceBindingSecretReferenceDie {
	return d.DieStamp(func(r *apiv1beta1.ServiceBindingSecretReference) {
		r.Name = v
	})
}
