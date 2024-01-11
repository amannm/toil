#!/usr/bin/env bash

EXEC_NAME="toil"
BUILD_DIR="build"

build() {
  local os="darwin"
  local arch="arm64"
  rm -rf "${BUILD_DIR}"
  mkdir -p "${BUILD_DIR}"
  GOOS="${os}" GOARCH="${arch}" go build -o "${BUILD_DIR}/${EXEC_NAME}"
}

"$@"