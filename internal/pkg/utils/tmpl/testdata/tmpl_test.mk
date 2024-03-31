# The binary to build.
BIN ?= {{ .Project.Bin }}

# This repo's root import path (under GOPATH).
PKG := {{ .Project.BuildPkg }}
VERSION_PKG={{ .Project.VersionPkg }}

ifeq ($(origin VERSION), undefined)
VERSION := $(shell git describe --tags --always --match='v*')
endif

GO_LDFLAGS += -X $(VERSION_PKG).Version=$(VERSION) \
	-X $(VERSION_PKG).GitCommit=$(GIT_COMMIT) \
	-X $(VERSION_PKG).GitTreeState=$(GIT_TREE_STATE) \
	-X $(VERSION_PKG).BuildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')