#!/bin/sh

OS=centos7
ARGS="-os=centos -from=centos:7 -maintainer signalsciences.com"
mkdir -p ${OS} 
cp testing/nginx.conf ${OS}
go run main.go ${ARGS} > ${OS}/Dockerfile
go run main.go ${ARGS} -style=docker-debug > ${OS}/Dockerfile.debug
go run main.go ${ARGS} -style=sh > ${OS}/build-lua-ngx.sh
chmod a+x ${OS}/build-lua-ngx.sh

OS=centos6
ARGS="-os=centos -from=centos:6 -maintainer signalsciences.com"
mkdir -p ${OS} 
cp testing/nginx.conf ${OS}
go run main.go ${ARGS} > ${OS}/Dockerfile
go run main.go ${ARGS} -style=docker-debug > ${OS}/Dockerfile.debug
go run main.go ${ARGS} -style=sh > ${OS}/build-lua-ngx.sh
chmod a+x ${OS}/build-lua-ngx.sh

OS=ubuntu1604
ARGS="-os=ubuntu -from=ubuntu:16.04 -maintainer signalsciences.com"
mkdir -p ${OS} 
cp testing/nginx.conf ${OS}
go run main.go ${ARGS} > ${OS}/Dockerfile
go run main.go ${ARGS} -style=docker-debug > ${OS}/Dockerfile.debug
go run main.go ${ARGS} -style=sh > ${OS}/build-lua-ngx.sh
chmod a+x ${OS}/build-lua-ngx.sh

OS=ubuntu1404
ARGS="-os=ubuntu -from=ubuntu:14.04 -maintainer signalsciences.com"
mkdir -p ${OS} 
cp testing/nginx.conf ${OS}
go run main.go ${ARGS} > ${OS}/Dockerfile
go run main.go ${ARGS} -style=docker-debug > ${OS}/Dockerfile.debug
go run main.go ${ARGS} -style=sh > ${OS}/build-lua-ngx.sh
chmod a+x ${OS}/build-lua-ngx.sh
