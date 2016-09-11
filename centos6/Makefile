CONTAINER=centos6-nginx-lua
build:
		docker build -t ${CONTAINER} .

console:
		docker run -it --rm ${CONTAINER} bash

history:
		docker history ${CONTAINER}
	
image:
		docker images ${CONTAINER}

.PHONY: build console
