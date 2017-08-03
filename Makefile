lint:
	gometalinter \
		--vendor \
		--vendored-linters \
		--deadline=60s \
		--disable-all \
		--enable=goimports \
		--enable=aligncheck \
		--enable=vetshadow \
		--enable=varcheck \
		--enable=structcheck \
		--enable=deadcode \
		--enable=ineffassign \
		--enable=unconvert \
		--enable=goconst \
		--enable=golint \
		--enable=gofmt \
		--enable=errcheck \
		--enable=misspell \
		./...

build:
	./generate.sh
