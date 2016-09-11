#!/bin/sh

OS=centos/7
ARGS="-os=centos -from=centos:7 -maintainer signalsciences.com"
mkdir -p ${OS} 
cp testing/nginx.conf ${OS}
go run main.go ${ARGS} > ${OS}/Dockerfile
go run main.go ${ARGS} -style=docker-debug > ${OS}/Dockerfile.debug
go run main.go ${ARGS} -style=sh > ${OS}/build-lua-ngx.sh
chmod a+x ${OS}/build-lua-ngx.sh

OS=centos/6
ARGS="-os=centos -from=centos:6 -maintainer signalsciences.com"
mkdir -p ${OS} 
cp testing/nginx.conf ${OS}
go run main.go ${ARGS} > ${OS}/Dockerfile
go run main.go ${ARGS} -style=docker-debug > ${OS}/Dockerfile.debug
go run main.go ${ARGS} -style=sh > ${OS}/build-lua-ngx.sh
chmod a+x ${OS}/build-lua-ngx.sh

OS=ubuntu/1604
ARGS="-os=ubuntu -from=ubuntu:16.04 -maintainer signalsciences.com"
mkdir -p ${OS} 
cp testing/nginx.conf ${OS}
go run main.go ${ARGS} > ${OS}/Dockerfile
go run main.go ${ARGS} -style=docker-debug > ${OS}/Dockerfile.debug
go run main.go ${ARGS} -style=sh > ${OS}/build-lua-ngx.sh
chmod a+x ${OS}/build-lua-ngx.sh

OS=ubuntu/1404
ARGS="-os=ubuntu -from=ubuntu:14.04 -maintainer signalsciences.com"
mkdir -p ${OS} 
cp testing/nginx.conf ${OS}
go run main.go ${ARGS} > ${OS}/Dockerfile
go run main.go ${ARGS} -style=docker-debug > ${OS}/Dockerfile.debug
go run main.go ${ARGS} -style=sh > ${OS}/build-lua-ngx.sh
chmod a+x ${OS}/build-lua-ngx.sh
