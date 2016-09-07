#!/bin/sh
set -ex
# FROM ubuntu:16.04
# MAINTAINER signalsciences.com

export NGINX_VERSION="${NGINX_VERSION:-1.10.1}"

export NGINX_LUA="${NGINX_LUA:-0.10.6}"

export NGINX_DEVEL="${NGINX_DEVEL:-0.3.0}"

export LUAJIT="${LUAJIT:-2.0.4}"
export top="${PWD}"
export tmpdir="/tmp/nginx"
export install_packages="make wget gcc autoconf automake libtool libc6-dev libc-dev libpcre3-dev zlib1g-dev libssl-dev pgp"
export modules_path="/etc/nginx/modules"
mkdir -p ${tmpdir} && cd ${tmpdir}
cp -f ${top}/nginx.conf /etc/nginx/nginx-helloworld.conf
apt-get update
apt-get install -y --no-install-recommends procps libpcre3 zlib1g openssl ca-certificates ${install_packages}
groupadd -r nginx
useradd -r -g nginx -s /sbin/nologin -d /var/cache/nginx -c "nginx user" nginx
wget -nv -O LuaJIT-${LUAJIT}.tar.gz http://luajit.org/download/LuaJIT-${LUAJIT}.tar.gz
wget -nv -O nginx-${NGINX_VERSION}.tar.gz http://nginx.org/download/nginx-${NGINX_VERSION}.tar.gz
wget -nv -O lua-nginx-module-${NGINX_LUA}.tar.gz https://github.com/openresty/lua-nginx-module/archive/v${NGINX_LUA}.tar.gz
wget -nv -O ngx_devel_kit-${NGINX_DEVEL}.tar.gz https://github.com/simpl/ngx_devel_kit/archive/v${NGINX_DEVEL}.tar.gz
tar -xf LuaJIT-${LUAJIT}.tar.gz
tar -xf lua-nginx-module-${NGINX_LUA}.tar.gz
tar -xf ngx_devel_kit-${NGINX_DEVEL}.tar.gz
tar -xf nginx-${NGINX_VERSION}.tar.gz
cd ${tmpdir}/LuaJIT-${LUAJIT} && make amalg BUILDMODE=static CC="gcc -fPIC"
cp ${tmpdir}/LuaJIT-${LUAJIT}/src/libluajit.a ${tmpdir}/LuaJIT-${LUAJIT}/src/libluajit-5.1.a
cd ${tmpdir}/nginx-${NGINX_VERSION} && \
    LUAJIT_LIB=${tmpdir}/LuaJIT-${LUAJIT}/src LUAJIT_INC=${tmpdir}/LuaJIT-${LUAJIT}/src ./configure \
    --prefix=/etc/nginx \
    --sbin-path=/usr/sbin/nginx \
    --modules-path=${modules_path} \
    --conf-path=/etc/nginx/nginx.conf \
    --error-log-path=/var/log/nginx/error.log \
    --http-log-path=/var/log/nginx/access.log \
    --pid-path=/var/run/nginx.pid \
    --lock-path=/var/run/nginx.lock \
    --http-client-body-temp-path=/var/cache/nginx/client_temp \
    --http-proxy-temp-path=/var/cache/nginx/proxy_temp \
    --http-fastcgi-temp-path=/var/cache/nginx/fastcgi_temp \
    --http-uwsgi-temp-path=/var/cache/nginx/uwsgi_temp \
    --http-scgi-temp-path=/var/cache/nginx/scgi_temp \
    --user=nginx \
    --group=nginx \
    --with-http_ssl_module \
    --with-http_realip_module \
    --with-http_addition_module \
    --with-http_sub_module \
    --with-http_dav_module \
    --with-http_flv_module \
    --with-http_mp4_module \
    --with-http_gunzip_module \
    --with-http_gzip_static_module \
    --with-http_random_index_module \
    --with-http_secure_link_module \
    --with-http_stub_status_module \
    --with-http_auth_request_module \
    --with-threads \
    --with-stream \
    --with-stream_ssl_module \
    --with-http_slice_module \
    --with-file-aio \
    --with-ipv6 \
    --with-http_v2_module \
    --with-cc-opt='-O2 -g -pipe -Wall -Wp,-D_FORTIFY_SOURCE=2 -fexceptions -fstack-protector-strong --param=ssp-buffer-size=4 -grecord-gcc-switches -m64 -mtune=generic' \
    --add-dynamic-module=../ngx_devel_kit-${NGINX_DEVEL} \
    --add-dynamic-module=../lua-nginx-module-${NGINX_LUA}
cd ${tmpdir}/nginx-${NGINX_VERSION} && make && make install
mkdir -p /var/cache/nginx/client_temp
rm -rf ${tmpdir}
apt-get purge -y ${install_packages}
apt-get autoremove -y
apt-get clean
