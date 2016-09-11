CONTAINER=centos7-nginx-lua
build:
		docker build -t ${CONTAINER} .

build-debug:
		docker build -t ${CONTAINER} -f Dockerfile.debug .
console:
		docker run -it --rm ${CONTAINER} bash

history:
		docker history ${CONTAINER}

image:
		docker images ${CONTAINER}

.PHONY: build build-debug console history image
