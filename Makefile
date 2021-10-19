include golang.mk
include lambda.mk
.DEFAULT_GOAL := test # override default goal set in library makefile

SHELL := /bin/bash
PKGS := $(shell go list ./... | grep -v /vendor | grep -v /tools)
CMD ?= dynamodb
REPONAME := $(notdir $(shell pwd))
PKG_MAIN := github.com/Clever/$(REPONAME)/cmd/$(CMD)
APP_NAME ?= $(REPONAME)

.PHONY: test build run $(PKGS) install_deps

$(eval $(call golang-version-check,1.16))

test: $(PKGS)

build: generate
	$(call lambda-build-go,$(PKG_MAIN),$(APP_NAME))

run: build
	echo "local run not supported yet, consider unit tests or deploying into dev"

$(PKGS): golang-test-all-deps
	$(call golang-test-all,$@)

generate:
	go generate ./cmd/$(CMD)

install_deps:
	go mod vendor
	go build -o bin/go-bindata ./vendor/github.com/kevinburke/go-bindata/go-bindata
