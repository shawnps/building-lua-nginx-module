CONTAINER=centos6-nginx-lua
build:
		docker build -t ${CONTAINER} .

console:
		docker run -it --rm ${CONTAINER} bash

.PHONY: build console
