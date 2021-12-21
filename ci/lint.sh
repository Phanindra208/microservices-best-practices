#!/bin/sh
set -u pipefail

GOBIN=$(pwd)/toolbin

GOIMPORTS=$GOBIN/goimports
GOLINT=$GOBIN/golint
STATICCHECK=$GOBIN/staticcheck

ALL_DIRS=$(go list -f '{{ .Dir }}' ./...)

FMT_RESULT=$($GOIMPORTS -l .)
if [ "$FMT_RESULT" != "" ]; then echo "unformatted files:"; echo "$FMT_RESULT"; exit 1; fi

set -eu pipefail

go vet $ALL_DIRS
$GOLINT -set_exit_status=1 $ALL_DIRS
$STATICCHECK $ALL_DIRS
