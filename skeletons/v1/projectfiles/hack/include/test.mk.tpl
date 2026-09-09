GINKGO := $(shell go env GOPATH)/bin/ginkgo
CLI ?= $(OUTPUT_DIR)/{{ .App.NormalizedName }}

.PHONY: test.cover
test.cover:
	@echo "===========> Run unit test"
	@go test -race -cover -coverprofile=$(REPO_ROOT)/coverage.out \
			-timeout=10m -short -v ./...

.PHONY: test.e2e
test.e2e: tools.verify.ginkgo
	@echo "===========> Run e2e test, CLI: $(CLI)"
	@$(GINKGO) -v $(REPO_ROOT)/test/e2e -- -cli=$(CLI)
