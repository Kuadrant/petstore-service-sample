SHELL := /bin/bash

IMG ?= quay.io/eastizle/petstore:1.0.0

all: build

build: download
	go build -o bin/petstore main.go

.PHONY: download
download:
	@echo Download go.mod dependencies
	@go mod download

# Run go fmt against code
.PHONY: fmt
fmt:
	go fmt ./...

# Run go vet against code
.PHONY: vet
vet:
	go vet ./...

.PHONY: build-image
build-image:
	docker build . -t ${IMG}

.PHONY: push-image
push-image:
	docker push ${IMG}

.PHONY: install
install:
	kubectl apply -f manifests/
	kubectl create configmap petstore-oas --from-file=./pkg/openapi/petstore.yaml

.PHONY: uninstall
uninstall:
	-kubectl delete -f manifests/
	-kubectl delete configmap petstore-oas

