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

## Startup

Currently, this does not provide or set up any init scripts, upstart config or systemd config as it was originally designed to run inside containers.

## Installation layout

The layout of nginx is the same as the official packages from [nginx.org](http://nginx.org/en/linux_packages.html).

## Testing

The "hello world" of nginx configs is loading in the [testing](/testing) directory and installed at `/etc/nginx/nginx-helloworld.conf`.

If `curl` is install, you can test the build by

```bash
# nginx -c /etc/nginx/nginx-helloworld.conf
# curl -sS http://127.0.0.1:80/lua_content
Hello,world
# nginx -s stop
```

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

One must add the following to their `nginx.conf` to use lua-ngx.

```
load_module modules/ndk_http_module.so; 
load_module modules/ngx_http_lua_module.so;
```

## The golang generator script

The bash scripts and the Dockerfiles are 99.9% identical.  The [main.go](/main.go) generator script takes a list of shell commands, and reformats them for use in either bash or in Docker.

You dont need to run it, since the generated files are checked in.  Feel free to use it for other purposes.

## Alpine Linux

[Alpine Linux](http://alpinelinux.org) provides lua-ngx-module in it's official repositories.  Search for [nginx-lua](http://pkgs.alpinelinux.org/packages?name=nginx-lua&branch=&repo=&arch=&maintainer=)

### Key Locations

RedHat / CentOS

```
nginx path prefix: "/etc/nginx"
nginx binary file: "/usr/sbin/nginx"
nginx modules path: "/usr/lib64/nginx/modules"
nginx configuration prefix: "/etc/nginx"
nginx configuration file: "/etc/nginx/nginx.conf"
nginx pid file: "/var/run/nginx.pid"
nginx error log file: "/var/log/nginx/error.log"
nginx http access log file: "/var/log/nginx/access.log"
nginx http client request body temporary files: "/var/cache/nginx/client_temp"
nginx http proxy temporary files: "/var/cache/nginx/proxy_temp"
nginx http fastcgi temporary files: "/var/cache/nginx/fastcgi_temp"
nginx http uwsgi temporary files: "/var/cache/nginx/uwsgi_temp"
nginx http scgi temporary files: "/var/cache/nginx/scgi_temp"
```

