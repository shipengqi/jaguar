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

.PHONY: release.verify
release.verify: tools.verify.git-chglog tools.verify.releaser

.PHONY: release.tag
release.tag: tools.verify.gsemver release.ensure-tag
	@git push origin `git describe --tags --abbrev=0`

.PHONY: release.ensure-tag
release.ensure-tag: tools.verify.gsemver
	@VERSION=$(VERSION) bash ./hack/ensure_tag.sh

## changelog: (Deprecated) generate changelogs.
.PHONY: release.changelog
release.changelog: tools.verify.git-chglog
	@gitversion=$(git describe --tags --abbrev=0)
	@VERSION=${VERSION:-gitversion}
	@REPO_ROOT=$(REPO_ROOT) bash ./hack/changelog.sh

.PHONY: release.run
release.run: release.verify
	@echo "===========> Releasing all build output"
	@gitversion=$(git describe --tags --abbrev=0)
	@VERSION=${VERSION:-gitversion}
	@GITHUB_TOKEN=$(GITHUB_TOKEN) \
	    RELEASE_FILE=$(REPO_ROOT)/CHANGELOG/CHANGELOG-$(VERSION).md \
		PUBLISH=$(PUBLISH) \
		GIT_TREE_STATE=$(GIT_TREE_STATE) \
		bash ./hack/release.sh
