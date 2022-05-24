# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Tools
CONTROLLER_GEN ?= go run -modfile hack/controller-gen/go.mod sigs.k8s.io/controller-tools/cmd/controller-gen
DIEGEN ?= go run -modfile hack/diegen/go.mod dies.dev/diegen
GOIMPORTS ?= go run -modfile hack/goimports/go.mod golang.org/x/tools/cmd/goimports
KAPP ?= go run -modfile hack/kapp/go.mod github.com/k14s/kapp/cmd/kapp
KO ?= go run -modfile hack/ko/go.mod github.com/google/ko
KUSTOMIZE ?= go run -modfile hack/kustomize/go.mod sigs.k8s.io/kustomize/kustomize/v4

KAPP_APP ?= servicebinding-runtime
KAPP_APP_NAMESPACE ?= default

# Setting SHELL to bash allows bash commands to be executed by recipes.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

.PHONY: all
all: build

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: manifests
manifests: ## Generate WebhookConfiguration, ClusterRole and CustomResourceDefinition objects.
	$(CONTROLLER_GEN) rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases

.PHONY: generate
generate: ## Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations.
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."
	$(DIEGEN) die:headerFile="hack/boilerplate.go.txt" paths="./..."

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...
	$(GOIMPORTS) --local github.com/scothis/servicebinding-runtime -w .

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: test
test: manifests generate fmt vet ## Run tests.
	go test ./... -coverprofile cover.out

##@ Build

.PHONY: build
build: generate fmt vet ## Build manager binary.
	go build -o bin/manager main.go

.PHONY: run
run: manifests generate fmt vet ## Run a controller from your host.
	go run ./main.go

##@ Deployment

.PHONY: install
install: manifests ## Install CRDs into the K8s cluster specified in ~/.kube/config.
	$(KAPP) deploy -a $(KAPP_APP) -n $(KAPP_APP_NAMESPACE) -f <($(KUSTOMIZE) build config/crd)

.PHONY: uninstall
uninstall: manifests ## Uninstall CRDs from the K8s cluster specified in ~/.kube/config.
	$(KAPP) delete -a $(KAPP_APP) -n $(KAPP_APP_NAMESPACE)

.PHONY: deploy
deploy: manifests ## Deploy controller to the K8s cluster specified in ~/.kube/config.
	$(KAPP) deploy -a $(KAPP_APP) -n $(KAPP_APP_NAMESPACE) -c -f config/kapp -f <($(KUSTOMIZE) build config/default | $(KO) resolve -f -)

.PHONY: undeploy
undeploy: ## Undeploy controller from the K8s cluster specified in ~/.kube/config.
	$(KAPP) delete -a $(KAPP_APP) -n $(KAPP_APP_NAMESPACE)
