#!/bin/bash

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


if [[ -z "${GO_LDFLAGS}" ]]; then
    echo "GO_LDFLAGS must be set"
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

export CGO_ENABLED=0

OUTPUT=${OUTPUT_DIR}/${BIN}
if [[ "${GOOS}" = "windows" ]]; then
  OUTPUT="${OUTPUT}.exe"
fi

CGO_ENABLED=0 go build \
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
