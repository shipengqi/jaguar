.DEFAULT_GOAL := all

# ==============================================================================
# Includes

include hack/include/common.mk # make sure include common.mk at the first include line
include hack/include/tools.mk
include hack/include/go.mk
include hack/include/test.mk
include hack/include/release.mk

all: modules lint test build

# ==============================================================================
# Usage

define USAGE_OPTIONS

选项:
  VERSION          编译进二进制的版本信息，默认从 gsemver 或 git 获取。
  PUBLISH          是否发布 release 到 Github，默认为 0。使用 make release 时有效。
  GITHUB_TOKEN     访问 Github 的 token，使用 make release 时有效。
  V                设置为 1 以启用详细构建信息输出，默认为 0。
  DEBUG            是否生成 debug 符号，默认为 0。
endef
export USAGE_OPTIONS

## --------------------------------------
## Build
## --------------------------------------

##@ build:

build: modules ## 编译二进制文件.
	@$(MAKE) go.build

modules: ## 添加缺失模块并移除未使用的模块.
	@go mod tidy

## --------------------------------------
## Release
## --------------------------------------

##@ release:

tag: ## 生成 release tag.
	@$(MAKE) release.tag

release: ## 发布一个版本.
	@$(MAKE) release.run

## --------------------------------------
## Testing
## --------------------------------------

##@ test:

test: ## 运行单元测试并获取覆盖率.
	@$(MAKE) test.cover

e2e: ## 运行 e2e 测试.
	@$(MAKE) test.e2e

## --------------------------------------
## Cleanup
## --------------------------------------

##@ clean:

clean: ## 清理构建产物.
	@$(MAKE) go.clean

## --------------------------------------
## Lint / Verification
## --------------------------------------

##@ lint and verify:

lint: ## 对 Go 代码执行静态检查.
	@$(MAKE) go.lint

## --------------------------------------
## Hack / Tools
## --------------------------------------

##@ hack/tools:

help: Makefile ## 打印 Makefile help 信息.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<TARGETS> <OPTIONS>\033[0m\n\n\033[35mTargets:\033[0m\n"} /^[0-9A-Za-z._-]+:.*?##/ { printf "  \033[36m%-45s\033[0m %s\n", $$1, $$2 } /^\$$\([0-9A-Za-z_-]+\):.*?##/ { gsub("_","-", $$1); printf "  \033[36m%-45s\033[0m %s\n", tolower(substr($$1, 3, length($$1)-7)), $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
	@printf "%s\n" "$$USAGE_OPTIONS"

.PHONY: all build modules tag release test e2e clean lint help
