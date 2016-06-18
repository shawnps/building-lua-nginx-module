# building-lua-nginx-module
scripts and Dockerfiles for building lua-nginx-module for nginx

Each directory contains three sample files:

| Name    | Function |
|---------|----------|
| Dockerfile | A all-in-one dockerfile that can be cut-n-paste into your own build system.  It makes the entire nginx + lua in one layer |
| Dockerfile.debug | Same as above but executes each command as a separate RUN statement.  This may be useful for debugging |
| install-lua-nginx.sh | the same functions but as a shell script (not docker) |

No images or containers are provided. They are designed to edited and copied
into your own build system as needed.

The `centos` recipes will work with CentOS or RHEL versions 6 and 7

The `ubuntu` recipes will work with Debian 7 and 8, and Ubuntu 12, 14 and 16.


