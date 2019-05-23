#!/bin/bash

cd "$(dirname "$0")"

set -e

buildArch=`uname -m`
case "${buildArch##*-}" in
	aarch64) ;;
  	x86_64) ;;
  	*) echo "Current architecture (${buildArch}) is not supported."; exit 2; ;;
esac

echo "==============================================="
echo " Image: ${IMAGE_TAG:=prod_vpp_agent}"
echo "==============================================="
echo " - dev image: ${DEV_IMG:=dev_vpp_agent}"
echo "==============================================="

docker build -f Dockerfile \
	--no-cache \
    --build-arg DEV_IMG=${DEV_IMG} \
	--tag ${IMAGE_TAG} \
	${DOCKER_BUILD_ARGS} .
