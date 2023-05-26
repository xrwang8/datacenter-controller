# Image URL to use all building/pushing image targets
IMG ?= registry.cnbita.com:5000/fusionhub/datacenter-controller:latest



.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: build
build:  fmt vet ## Build  binary.
	go build -o bin/datacenter main.go

.PHONY: run
run:   fmt vet
	go run ./main.go

.PHONY: docker-build
docker-build:  ## Build docker image with the manager.
	docker build -t ${IMG} .


.PHONY: docker-push
docker-push: ## Push docker image with the manager.
	docker push ${IMG}
