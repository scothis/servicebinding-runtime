# Service Binding Runtime <!-- omit in toc -->

Experimental implementation of the [ServiceBinding.io](https://servicebinding.io) [1.0 spec](https://servicebinding.io/spec/core/1.0.0/) using [reconciler-runtime](https://github.com/vmware-labs/reconciler-runtime/). The full specification is implemented, please open an issue for any discrepancies.

- [Architecture](#architecture)
  - [Controller](#controller)
  - [Webhooks](#webhooks)
- [Supported Services](#supported-services)
- [Supported Workloads](#supported-workloads)
- [Getting Started](#getting-started)
  - [Running on the cluster](#running-on-the-cluster)
  - [Undeploy controller](#undeploy-controller)
- [Contributing](#contributing)
  - [Test It Out](#test-it-out)
  - [Modifying the API definitions](#modifying-the-api-definitions)
- [Acknowledgements](#acknowledgements)
- [License](#license)

## Architecture

The [Service Binding for Kubernetes Specification](https://servicebinding.io/spec/core/1.0.0/) defines the shape of [Provisioned Services](https://servicebinding.io/spec/core/1.0.0/#provisioned-service), and how the `Secret` is [projected into a workload](https://servicebinding.io/spec/core/1.0.0/#workload-projection). The spec says less (intentionally) about how this happens.

Both a controller and mutating admission webhook are used to project a `Secret` defined by the service referenced by the `ServiceBinding` resource into the workloads referenced. The controller is used to process `ServiceBinding`s by resolving services, projecting workloads and updating the status. The webhook is used to prevent removal of the workload projection, projecting workload on create, and a notification trigger for `ServiceBinding`s the controller should process.

The apis, resolver and projector packages are defined by the reference implementation and reused here with slight modifications. The bulk of the work to bind a service to a workload is encapsulated with these packages. The output from the projector is deterministic and idempotent. The order that service bindings are applied to, or removed from, a workload does not matter. If a workload is bound and then unbound, the only trace will be the `SERVICE_BINDING_ROOT` environment variable.

There are a limited number of resources that maintain an informer cache within the manager:
- `ServiceBinding`
- `ClusterWorkloadResourceMapping`
- `MutatingWebhookConfiguration`
- `ValidatingWebhookConfiguration`

### Controller

When a `ServiceBinding` is created, updated or deleted the controller processes the resource. It will:
- resolve the referenced service resource, looking at it's `.spec.binding.name` for the name of the Secret to bind
- reflect the discovered `Secret` name onto the `ServiceBinding`'s `.status.binding.name`
- the `ServiceAvailable` condition is updated on the `ServiceBinding`
- the references workloads are resolved (either by name or selector)
- a `ClusterWorkloadResourceMapping` is resolved for the apiVersion/kind of the workload (or a default value for a PodSpecable workload is used)
- the resolved `Secret` name is projected into the workload
- the `Ready` condition is updated on the `ServiceBinding`

### Webhooks

In addition to that main flow, a `MutatingWebhookConfiguration` and `ValidationWebhookConfiguration` are updated:
- all `ServiceBinding`s in the cluster are resolved
- the rules for a MutatingWebhookConfiguration are updated based on the set of all workload group-kinds referenced
- the rules for a ValidatingWebhookConfiguration are updated based on the set of all workload and service group-kinds referenced

The `MutatingWebhookConfiguration` is used to intercept create and update requests for workloads:
- all `ServiceBinding`s targeting the workload are resolved
- a `ClusterWorkloadResourceMapping` is resolved for the apiVersion/kind of the workload (or a default value for a PodSpecable workload is used)
- for each `ServiceBinding` the resolved `Secret` name is projected into the workload
- the delta between the original resource and the projected resource is returned with the webhook response as a patch

The `ValidationWebhookConfiguration` is used as an alternative to watching the API Server directly for these types and keeping an informer cache. When a webhook request is received, the `ServiceBinding`s that reference that resource as a workload or service are resolved and enqueued for the controller to process.

No blocking work is performed within the webhooks.

## Supported Services

Kubernetes defines no provisioned services by default, however, `Secret`s may be [directly referenced](https://servicebinding.io/spec/core/1.0.0/#direct-secret-reference).

Additional services can be supported dynamically by [defining a `ClusterRole`](https://servicebinding.io/spec/core/1.0.0/#considerations-for-role-based-access-control-rbac).

## Supported Workloads

Support for the built-in k8s workload resource is pre-configured including:
- apps `DaemonSet`
- apps `Deployment`
- apps `ReplicaSet`
- apps `StatefulSet`
- batch `CronJob` (also includes a `ClusterResourceMapping`)
- batch `Job` (since Jobs are immutable, the ServiceBinding must be defined and service resolved before the job is created)
- core `ReplicationController`

Additional workloads can be supported dynamically by [defining a `ClusterRole`](https://servicebinding.io/spec/core/1.0.0/#considerations-for-role-based-access-control-rbac-1) and if not PodSpecable, a [`ClusterWorkloadResourceMapping`](https://servicebinding.io/spec/core/1.0.0/#workload-resource-mapping).


## Getting Started
You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Running on the cluster
1. Define where to publish images:

   ```sh
   export KO_DOCKER_REPO=<a-repository-you-can-write-to>
   ```

   For kind, a registry is not required:

   ```sh
   export KO_DOCKER_REPO=kind.local/servicebinding
   ```
	
1. Build and Deploy the controller to the cluster:

   ```sh
   make deploy
   ```

### Undeploy controller
Undeploy the controller to the cluster:

```sh
make undeploy
```

## Contributing

### Test It Out

Run the unit tests:

```sh
make test
```

### Modifying the API definitions
If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## Acknowledgements

The scope of this work is defined by the Service Binding for Kubernetes project. Without our friends from IBM and RedHat, this domain this controller operates within would not exist.

Large chunks of this code base are from the upstream [reference implementation](https://github.com/servicebinding/service-binding-controller) for the service binding specification, include the apis, projector and resolver packages. These sources retain the upstream copyright headers

## License

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

