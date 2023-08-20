#!/bin/bash

if [[ -z "${GO_LDFLAGS}" ]]; then
    echo "GO_LDFLAGS must be set"
    exit 1
fi

echo "Build {{ .Project.Name }} FAILED"
