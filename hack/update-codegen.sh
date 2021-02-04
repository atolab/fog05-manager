#!/usr/bin/env bash

# ROOT_PACKAGE :: the package (relative to $GOPATH/src) that is the target for code generation
ROOT_PACKAGE="github.com/atolab/fog05-manager"
# CUSTOM_RESOURCE_NAME :: the name of the custom resource that we're generating client code for
CUSTOM_RESOURCE_NAME="fdu"
# CUSTOM_RESOURCE_VERSION :: the version of the resource
CUSTOM_RESOURCE_VERSION="v1"


set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
OUTPUT="$(dirname "${BASH_SOURCE[0]}")/.."
echo "$OUTPUT"


bash vendor/k8s.io/code-generator/generate-groups.sh all "$ROOT_PACKAGE/pkg/client" "$ROOT_PACKAGE/pkg/apis" "$CUSTOM_RESOURCE_NAME:$CUSTOM_RESOURCE_VERSION" --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt

rm -rf $SCRIPT_ROOT/pkg/client
mv $SCRIPT_ROOT/$ROOT_PACKAGE/pkg/apis/fdu/v1/zz_generated.deepcopy.go $SCRIPT_ROOT/pkg/apis/fdu/v1/zz_generated.deepcopy.go
mv $SCRIPT_ROOT/$ROOT_PACKAGE/pkg/client $SCRIPT_ROOT/pkg/client
rm -rf $SCRIPT_ROOT/github.com