#!/bin/bash

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

set -o errexit
set -o nounset

if [[ -z "${PKG}" ]]; then
    echo "PKG must be set"
    exit 1
fi
if [[ -z "${BIN}" ]]; then
    echo "BIN must be set"
    exit 1
fi
if [[ -z "${GOOS}" ]]; then
    echo "GOOS must be set"
    exit 1
fi
if [[ -z "${GOARCH}" ]]; then
    echo "GOARCH must be set"
    exit 1
fi

if [[ -z "${GO_LDFLAGS}" ]]; then
    echo "GO_LDFLAGS must be set"
    exit 1
fi

if [[ -z "${VERSION}" ]]; then
    echo "VERSION must be set"
    exit 1
fi

if [[ -z "${GIT_COMMIT}" ]]; then
    echo "GIT_COMMIT must be set"
    exit 1
fi

if [[ -z "${GIT_TREE_STATE}" ]]; then
    echo "GIT_TREE_STATE must be set"
    exit 1
fi

if [[ -z "${OUTPUT_DIR}" ]]; then
    echo "OUTPUT_DIR must be set"
    exit 1
fi

GCFLAGS=""
if [[ ${DEBUG:-} = "1" ]]; then
    GCFLAGS="all=-N -l"
fi

if [[ -z "${BUILD_TIME}" ]]; then
    BUILD_TIME=$(date -u +'%Y-%m-%dT%H:%M:%SZ')
fi

export CGO_ENABLED=0

OUTPUT=${OUTPUT_DIR}/${BIN}
if [[ "${GOOS}" = "windows" ]]; then
  OUTPUT="${OUTPUT}.exe"
fi

go build \
    -o ${OUTPUT} \
    -gcflags "${GCFLAGS}" \
    -ldflags "${GO_LDFLAGS}" \
    ${PKG}

if [[ "$?" -eq 0 ]];then
    echo "Build ${OUTPUT} SUCCESS"
else
    echo "Build ${OUTPUT} FAILED"
    exit 1
fi
