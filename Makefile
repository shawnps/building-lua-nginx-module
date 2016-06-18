
build:
	mkdir -p centos ubuntu
	cp testing/nginx.conf centos/
	cp testing/nginx.conf ubuntu/
	go run main.go -os centos -from centos:7 -maintainer signalsciences.com > centos/Dockerfile
	go run main.go -os=centos -from=centos:7 -maintainer signalsciences.com -style=docker-debug > centos/Dockerfile.debug
	go run main.go -os=centos -from=centos:7 -maintainer signalsciences.com -style=sh > centos/build-lua-ngx.sh
	chmod a+x centos/build-lua-ngx.sh
	go run main.go -os=ubuntu -from=ubuntu:16.04 -maintainer signalsciences.com > ubuntu/Dockerfile
	go run main.go -os=ubuntu -from=ubuntu:16.04 -maintainer signalsciences.com -style=docker-debug > ubuntu/Dockerfile.debug
	go run main.go -os=ubuntu -from=ubuntu:16.04 -maintainer signalsciences.com -style=sh > ubuntu/build-lua-ngx.sh
	chmod a+x ubuntu/build-lua-ngx.sh
