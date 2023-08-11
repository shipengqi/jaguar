# The binary to build.
BIN ?= jaguar

# This repo's root import path (under GOPATH).
PKG := github.com/shipengqi/jaguar/cmd/jaguar
VERSION_PKG=github.com/shipengqi/component-base/pkg/version

ifeq ($(origin VERSION), undefined)
VERSION := $(shell git describe --tags --always --match='v*')
endif

ifeq ($(origin REPO_ROOT),undefined)
REPO_ROOT := $(shell git rev-parse --show-toplevel)
endif

# set git commit and tree state
GIT_COMMIT = $(shell git rev-parse HEAD)
ifneq ($(shell git status --porcelain 2> /dev/null),)
	GIT_TREE_STATE ?= dirty
else
	GIT_TREE_STATE ?= clean
endif

ARCH ?= $(shell go env GOOS)-$(shell go env GOARCH)
platform_temp = $(subst -, ,$(ARCH))
GOOS = $(word 1, $(platform_temp))
GOARCH = $(word 2, $(platform_temp))

ifeq ($(origin OUTPUT_DIR),undefined)
OUTPUT_DIR := $(REPO_ROOT)/_output/$(GOOS)/$(GOARCH)/bin
$(shell mkdir -p $(OUTPUT_DIR))
endif

# Specify tools.
BUILD_TOOLS ?= gsemver golangci-lint releaser ginkgo

# Makefile settings
# The --no-print-directory option of 'make' tells 'make' not to print
# the message about entering and leaving the working directory.
ifndef V
MAKEFLAGS += --no-print-directory
endif

ifeq ($(origin PUBLISH),undefined)
PUBLISH := 0
endif


GO_LDFLAGS += -X $(VERSION_PKG).GitVersion=$(VERSION) \
	-X $(VERSION_PKG).GitCommit=$(GIT_COMMIT) \
	-X $(VERSION_PKG).GitTreeState=$(GIT_TREE_STATE) \
	-X $(VERSION_PKG).BuildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')