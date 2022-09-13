B=$(shell git rev-parse --abbrev-ref HEAD)
BRANCH=$(subst /,-,$(B))
GITREV=$(shell git describe --abbrev=7 --always --tags)
REV=$(GITREV)-$(BRANCH)-$(shell date +%Y%m%d-%H:%M:%S)

docker:
	docker build -t umputun/echo-http .

dist:
	- @mkdir -p dist
	docker build -f Dockerfile.artifacts -t echo-http.bin .
	- @docker rm -f echo-http.bin 2>/dev/null || exit 0
	docker run -d --name=echo-http.bin echo-http.bin
	docker cp echo-http.bin:/artifacts dist/
	docker rm -f echo-http.bin

build: info
	- GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-X main.revision=$(REV) -s -w" -o ./dist/echo-http

info:
	- @echo "revision $(REV)"

.PHONY: dist docker race_test bin info
