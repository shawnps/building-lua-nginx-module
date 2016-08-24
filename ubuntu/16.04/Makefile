CONTAINER=ubuntu1604-nginx-lua
build:
		docker build -t ${CONTAINER} .

console:
		docker run -it --rm ${CONTAINER} bash

.PHONY: build
.PHONY: console
