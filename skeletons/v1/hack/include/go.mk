# Copyright (c) 2022 PengQi Shi
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

GO_SUPPORTED_VERSIONS ?= 1.17|1.18|1.19|1.20

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
	@echo "===========> Building binary: $(OUTPUT_DIR)/$(BIN)"
	@GOOS=$(GOOS) GOARCH=$(GOARCH) VERSION=$(VERSION) \
		PKG=$(PKG) BIN=$(BIN) \
		GIT_COMMIT=$(GIT_COMMIT) GIT_TREE_STATE=$(GIT_TREE_STATE) \
		OUTPUT_DIR=$(OUTPUT_DIR) \
		BUILD_TIME="" \
		bash ./hack/build.sh

.PHONY: go.lint
go.lint: tools.verify.golangci-lint
	@echo "===========> Run golangci-lint to lint source codes"
	@golangci-lint run -c $(REPO_ROOT)/.golangci.yaml $(REPO_ROOT)/...

.PHONY: go.clean
go.clean:
	@echo "===========> Cleaning all build output"
	@-rm -vrf $(OUTPUT_DIR)