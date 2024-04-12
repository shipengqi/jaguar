#!/bin/bash

set -o errexit
set -o nounset

RELEASER=$(go env GOPATH)/bin/goreleaser

# $PUBLISH must explicitly be set to '1' for goreleaser
# to publish the release to GitHub.
if [[ "${PUBLISH:-}" != "1" ]]; then
    echo "Not set to publish"
    ${RELEASER} release \
        --clean \
        --snapshot \
        --skip-publish
else
    if [[ -z "${GITHUB_TOKEN}" ]]; then
        echo "GITHUB_TOKEN must be set"
        exit 1
    fi

    echo "Getting ready to publish"
    ${RELEASER} release \
    --clean
fi
