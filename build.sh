#!/usr/bin/env bash

BGO_ROOT=$(git rev-parse --show-toplevel)
APP="httpbee"
VERSION=$(git describe --tags --always)
COMMIT_HASH=$(git rev-parse --short HEAD 2>/dev/null)
DATE=$(date "+%Y-%m-%d")
PROJECT="github.com/haozibi/${APP}"
BGO_BUILD_PLATFORMS="linux windows darwin"
BGO_BUILD_ARCHS="amd64 386 arm arm64"

echo ${BGO_ROOT}

if [[ "$(pwd)" != "${BGO_ROOT}" ]]; then
  echo "you are not in the root of the repo" 1>&2
  echo "please cd to ${BGO_ROOT} before running this script" 1>&2
  exit 1
fi

GO_BUILD_CMD="go build -a"
# GO_BUILD_CMD="go build -a -installsuffix"

GO_BUILD_LDFLAGS="-s -w -X ${PROJECT}/app.CommitHash=${COMMIT_HASH} -X ${PROJECT}/app.BuildTime=${DATE} -X ${PROJECT}/app.BuildVersion=${VERSION} -X ${PROJECT}/app.BuildAppName=${APP}"

mkdir -p "${BGO_ROOT}/release"

for OS in ${BGO_BUILD_PLATFORMS[@]}; do
  for ARCH in ${BGO_BUILD_ARCHS[@]}; do
    NAME="${APP}-${VERSION}-${OS}-${ARCH}"
    if [[ "${OS}" == "windows" ]]; then
      NAME="${NAME}.exe"
    fi

    # Enable CGO if building for OS X on OS X; see
    # https://github.com/golang/dep/issues/1838 for details.
    if [[ "${OS}" == "darwin" && "${BUILD_PLATFORM}" == "darwin" ]]; then
      CGO_ENABLED=1
    else
      CGO_ENABLED=0
    fi
    if [[ "${ARCH}" == "ppc64" || "${ARCH}" == "ppc64le" || "${ARCH}" == "s390x" || "${ARCH}" == "arm" || "${ARCH}" == "arm64" ]] && [[ "${OS}" != "linux" ]]; then
        # ppc64, ppc64le, s390x, arm and arm64 are only supported on Linux.
        echo "Building for ${OS}/${ARCH} not supported."
    else
        echo "Building for ${OS}/${ARCH} with CGO_ENABLED=${CGO_ENABLED}"
        GOARCH=${ARCH} GOOS=${OS} CGO_ENABLED=${CGO_ENABLED} ${GO_BUILD_CMD} -ldflags "${GO_BUILD_LDFLAGS}"\
            -o "${BGO_ROOT}/release/${NAME}" main.go
        pushd "${BGO_ROOT}/release" > /dev/null
        shasum -a 256 "${NAME}"  >> "${APP}-${VERSION}.sha256"
        popd > /dev/null
    fi
  done
done