package main

// This generates various scripts to automate building the
// lua-nginx-module

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

var centos_installs = "wget gcc autoconf automake libtool pcre-devel openssl-devel file which"

var centos_header = []string{
	`yum install -y ${install_packages} pcre openssl`,
	`groupadd -f -r nginx`,
	`useradd -r -g nginx -s /sbin/nologin -d /var/cache/nginx -c "nginx user" nginx`,
}

var centos_footer = []string{
	`yum remove -y ${install_packages}`,
	`yum clean all`,
	`rm -rf /var/cache/yum/* ${tmpdir}`,
}

var ubuntu_installs = "make wget gcc autoconf automake libtool libc6-dev libc-dev libpcre3-dev zlib1g-dev libssl-dev pgp"

var ubuntu_header = []string{
	`apt-get update`,
	`apt-get install -y --no-install-recommends procps libpcre3 zlib1g openssl ca-certificates ${install_packages}`,
	`groupadd -r nginx`,
	`useradd -r -g nginx -s /sbin/nologin -d /var/cache/nginx -c "nginx user" nginx`,
}

var ubuntu_footer = []string{
	`rm -rf ${tmpdir}`,
	`apt-get purge -y ${install_packages}`,
	`apt-get autoremove -y`,
	`apt-get clean`,
}

var build_nginx = []string{
	`wget -nv -O ${tmpdir}/checksec https://raw.githubusercontent.com/slimm609/checksec.sh/master/checksec`,
	`wget -nv -O LuaJIT-${LUAJIT}.tar.gz http://luajit.org/download/LuaJIT-${LUAJIT}.tar.gz`,
	`wget -nv -O nginx-${NGINX_VERSION}.tar.gz http://nginx.org/download/nginx-${NGINX_VERSION}.tar.gz`,
	`wget -nv -O lua-nginx-module-${NGINX_LUA}.tar.gz https://github.com/openresty/lua-nginx-module/archive/v${NGINX_LUA}.tar.gz`,
	`wget -nv -O ngx_devel_kit-${NGINX_DEVEL}.tar.gz https://github.com/simpl/ngx_devel_kit/archive/v${NGINX_DEVEL}.tar.gz`,
	`tar -xf LuaJIT-${LUAJIT}.tar.gz`,
	`tar -xf lua-nginx-module-${NGINX_LUA}.tar.gz`,
	`tar -xf ngx_devel_kit-${NGINX_DEVEL}.tar.gz`,
	`tar -xf nginx-${NGINX_VERSION}.tar.gz`,
	`cd ${tmpdir}/LuaJIT-${LUAJIT} && make amalg BUILDMODE=static CC="gcc -fPIC"`,
	`cp ${tmpdir}/LuaJIT-${LUAJIT}/src/libluajit.a ${tmpdir}/LuaJIT-${LUAJIT}/src/libluajit-5.1.a`,
	`cd ${tmpdir}/nginx-${NGINX_VERSION} && \
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
    --with-cc-opt='-O2 -g -pipe -Wall -fexceptions -m64 -mtune=generic ${cflag_extra}' \
    --with-ld-opt='-Wl,-Bsymbolic-functions -Wl,-z,relro -Wl,-z,now' \
    --add-dynamic-module=../ngx_devel_kit-${NGINX_DEVEL} \
    --add-dynamic-module=../lua-nginx-module-${NGINX_LUA}`,
	`cd ${tmpdir}/nginx-${NGINX_VERSION} && make && make install`,
	`/bin/bash -f ${tmpdir}/checksec -f /usr/sbin/nginx`,
	`/bin/bash -f ${tmpdir}/checksec -f ${modules_path}/ndk_http_module.so`,
	`/bin/bash -f ${tmpdir}/checksec -f ${modules_path}/ngx_http_lua_module.so`,
	`mkdir -p /var/cache/nginx/client_temp`,
}

var nginx_test = []string{
	`nginx -c /etc/nginx/nginx-helloworld.conf`,
	`curl -sS http://127.0.0.1/lua_content`,
	`nginx -s stop`,
}

type Generator interface {
	Copy(string, string) string
	From(string) string
	Maintainer(string) string
	Arg(string, string) string
	Env([][2]string) string
	SetEnv(string, string) string
	Run([]string) string
	Workdir(string) string
}

type DockerGenerator struct {
	Debug bool
}

func (d *DockerGenerator) Copy(src, dest string) string {
	return fmt.Sprintf("COPY %s %s", src, dest)
}

func (d *DockerGenerator) From(arg string) string {
	return "FROM " + arg
}

func (d *DockerGenerator) Maintainer(arg string) string {
	return "MAINTAINER " + arg
}

func (d *DockerGenerator) Arg(key, value string) string {
	if value == "" {
		return fmt.Sprintf("ARG %s", key)
	}
	return fmt.Sprintf("ARG %s=%q", key, value)
}

func (d *DockerGenerator) SetEnv(key, value string) string {
	return fmt.Sprintf("ENV %s=%q", key, value)
}

func (d *DockerGenerator) Env(args [][2]string) string {
	lines := make([]string, 0, len(args))
	for _, kv := range args {
		lines = append(lines, fmt.Sprintf("  %s=%s", kv[0], kv[1]))
	}
	return "ENV \\\n" + strings.Join(lines, "  \\\n")
}

func (d *DockerGenerator) Run(cmds []string) string {
	if !d.Debug {
		lines := []string{"RUN set -ex"}
		for _, line := range cmds {
			line = "  && " + strings.TrimSpace(line)
			lines = append(lines, line)
		}
		return strings.Join(lines, " \\\n")
	}

	lines := make([]string, 0, len(cmds))
	for _, line := range cmds {
		lines = append(lines, "RUN "+line)
	}
	return strings.Join(lines, "\n")

}

func (d *DockerGenerator) Workdir(dir string) string {
	return "WORKDIR " + dir
}

type ShellGenerator struct {
}

func (d *ShellGenerator) Copy(src, dest string) string {
	return fmt.Sprintf("cp -f %s %s", src, dest)
}

func (d *ShellGenerator) From(arg string) string {
	return "#!/bin/sh\nset -ex\n# FROM " + arg
}
func (d *ShellGenerator) Maintainer(arg string) string {
	return "# MAINTAINER " + arg
}
func (d *ShellGenerator) Arg(key, value string) string {
	if value == "" {
		return ""
	}
	return fmt.Sprintf("export %s=%q", key, value)
}

func (d *ShellGenerator) SetEnv(key, value string) string {
	return fmt.Sprintf("export %s=%q", key, value)
}

func (d *ShellGenerator) Env(args [][2]string) string {
	lines := make([]string, 0, len(args))
	for _, kv := range args {
		lines = append(lines, fmt.Sprintf("export %s=%q", kv[0], kv[1]))
	}
	return strings.Join(lines, "\n")
}

func (d *ShellGenerator) Run(cmds []string) string {
	lines := make([]string, 0, len(cmds))
	for _, line := range cmds {
		lines = append(lines, strings.TrimSpace(line))
	}
	return strings.Join(lines, "\n")
}

func (d *ShellGenerator) Workdir(dir string) string {
	return fmt.Sprintf("mkdir -p %s && cd %s", dir, dir)
}

func mergeLines(groups ...[]string) []string {
	lines := []string{}
	for _, g := range groups {
		for _, line := range g {
			lines = append(lines, strings.TrimSpace(line))
		}
	}
	return lines
}

func main() {
	shStyle := flag.String("style", "docker", "shell output style [sh|docker|docker-debug")
	argFrom := flag.String("from", "centos:7", "Docker FROM image")
	argOS := flag.String("os", "centos", "OS type centos or ubuntu")
	argMaintainer := flag.String("maintainer", "unknown", "Docker maintainer")

	flag.Parse()

	var gen Generator
	switch *shStyle {
	case "sh":
		gen = &ShellGenerator{}
	case "docker-debug":
		gen = &DockerGenerator{
			Debug: true,
		}
	case "docker":
		gen = &DockerGenerator{}
	default:
		log.Fatalf("-style must be one of sh,docker,docker-debug")
	}

	env := [][2]string{
		{"NGINX_VERSION", "1.10.1"},
		{"NGINX_LUA", "0.10.6"},
		{"NGINX_DEVEL", "0.3.0"},
		{"LUAJIT", "2.0.4"},
	}
	var cmds []string
	var installs string
	var modulesPath string
	var cflagSecurity string

	switch *argOS {
	case "centos", "rhel", "redhat":
		cmds = mergeLines(centos_header, build_nginx, centos_footer)
		installs = centos_installs
		modulesPath = "/usr/lib64/nginx/modules"
	case "debian", "ubuntu":
		cmds = mergeLines(ubuntu_header, build_nginx, ubuntu_footer)
		installs = ubuntu_installs
		modulesPath = "/etc/nginx/modules"
	default:
		log.Fatalf("Unknown OS type: should be centos or ubuntu")
	}

	// hacks around various old compilers.  We are explicity doing the replacement
	// here and not in the dockerfile so `nginx -V` shows the flags instead of ${cflag_extra}
	switch *argFrom {
	case "centos:6":
		// needed for gcc < 4.9
		cflagSecurity = "-Wp,-D_FORTIFY_SOURCE=2 -fstack-protector --param ssp-buffer-size=4"
	default:
		// gcc > 4.9
		cflagSecurity = "-Wp,-D_FORTIFY_SOURCE=2 -fstack-protector-strong"
	}
	for i, line := range cmds {
		cmds[i] = strings.Replace(line, "${cflag_extra}", cflagSecurity, -1)
	}

	fmt.Printf("%s\n", gen.From(*argFrom))
	fmt.Printf("%s\n", gen.Maintainer(*argMaintainer))

	for _, arg := range env {
		k, v := arg[0], arg[1]
		fmt.Printf("%s\n", gen.Arg(k, ""))
		fmt.Printf("%s\n", gen.SetEnv(k, fmt.Sprintf("${%s:-%s}", k, v)))
	}
	fmt.Printf("%s\n", gen.Arg("top", "${PWD}"))
	fmt.Printf("%s\n", gen.Arg("tmpdir", "/tmp/nginx"))
	fmt.Printf("%s\n", gen.Arg("install_packages", installs))
	fmt.Printf("%s\n", gen.Arg("modules_path", modulesPath))
	fmt.Printf("%s\n", gen.Workdir("${tmpdir}"))
	fmt.Printf("%s\n", gen.Copy("${top}/nginx.conf", "/etc/nginx/nginx-helloworld.conf"))
	fmt.Printf("%s\n", gen.Run(cmds))
	//fmt.Printf("%s\n", gen.Run(nginx_test))
}
