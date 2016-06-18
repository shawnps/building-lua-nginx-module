# building-lua-nginx-module
scripts and Dockerfiles for building lua-nginx-module for nginx

Each directory contains three sample files to aid in building [nginx](http://nginx.org) + [lua-ngx-module](https://github.com/openresty/lua-nginx-module) from source.  It takes about 2 minutes to build everything.

No images or containers are provided. They are designed to edited and copied
into your own build system as needed.

The [centos](/centos) recipes will work with CentOS or RHEL versions 6 and 7

The [ubuntu](/ubuntu) recipes will work with Debian 7 and 8, and Ubuntu 12, 14 and 16.


| Name    | Function |
|---------|----------|
| [install-lua-nginx.sh](/centos/install-lua-nginx.sh) | A shell script to download, build and install nginx + lua |
| [Dockerfile](/centos/Dockerfile) | An all-in-one dockerfile that can be cut-n-paste into your own build system.  It builds the entire nginx + lua in one layer to make the smallest possible container |
| [Dockerfile.debug](/centos/Dockerfile.debug) | Same as above but executes each command as a separate RUN statement.  This may be useful for debugging. |

## Installation layout

The layout of nginx is the same as the official packages from [nginx.org](http://nginx.org/en/linux_packages.html).

## Dependencies

All dependencies, include the version of nginx itself, are configurable by adjusting the environment variables defined in the beginning of the script.

```
ENV \
  NGINX_VERSION=1.10.1  \
  NGINX_LUA=0.10.5  \
  NGINX_DEVEL=0.3.0  \
  LUAJIT=2.0.4
```

## Static vs. Dynamic Modules

The current version makes dynamic modules for ngx-lua. 

## Alpine Linux

[Alpine Linux](http://alpinelinux.org) provides lua-ngx-module in it's official repositories.  Search for [nginx-lua](http://pkgs.alpinelinux.org/packages?name=nginx-lua&branch=&repo=&arch=&maintainer=)

