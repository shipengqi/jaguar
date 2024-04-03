.PHONY: release.verify
release.verify: tools.verify.releaser

.PHONY: release.tag
release.tag: tools.verify.gsemver release.ensure-tag
	@git push origin `git describe --tags --abbrev=0`

.PHONY: release.ensure-tag
release.ensure-tag: tools.verify.gsemver
	@VERSION=$(VERSION) bash $(REPO_ROOT)/hack/ensure_tag.sh

.PHONY: release.run
release.run: release.verify
	@echo "===========> Releasing all build output"
	@gitversion=$(git describe --tags --abbrev=0)
	@VERSION=${VERSION:-gitversion}
	@GITHUB_TOKEN=$(GITHUB_TOKEN) \
		PUBLISH=$(PUBLISH) \
		GO_LDFLAGS="$(GO_LDFLAGS)" \
		bash $(REPO_ROOT)/hack/release.sh
