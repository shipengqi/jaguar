GO_SUPPORTED_VERSIONS ?= 1.18|1.19|1.20|1.21|1.22

.PHONY: go.build.verify
go.build.verify:
ifneq ($(shell go version | grep -q -E '\bgo($(GO_SUPPORTED_VERSIONS))\b' && echo 0 || echo 1), 0)
	$(error unsupported go version. Please install one of the following supported version: '$(GO_SUPPORTED_VERSIONS)')
endif

.PHONY: go.build.dirs
go.build.dirs:
	@mkdir -p $(OUTPUT_DIR)

.PHONY: go.build
go.build: go.build.verify go.build.dirs
	@echo "===========> Building: $(OUTPUT_DIR)/$(BIN)"
	@GOOS=$(GOOS) \
		PKG=$(PKG) BIN=$(BIN) \
		OUTPUT_DIR=$(OUTPUT_DIR) \
		GO_LDFLAGS="$(GO_LDFLAGS)" \
		bash $(REPO_ROOT)/hack/build.sh

.PHONY: go.lint
go.lint: tools.verify.golangci-lint
	@echo "===========> Run golangci-lint to lint source codes"
	@golangci-lint run -c $(REPO_ROOT)/.golangci.yaml $(REPO_ROOT)/...

# `-` indicates that ignore the command error
# `-rm -vrf $(OUTPUT_DIR)` ignore if rm command execute error.
.PHONY: go.clean
go.clean:
	@echo "===========> Cleaning all build output"
	@-rm -vrf $(OUTPUT_DIR)
