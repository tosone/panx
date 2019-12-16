BuildStamp = $(shell date '+%Y-%m-%d_%H:%M:%S')
GitHash    = $(shell git rev-parse HEAD)
Version    = $(shell git describe --abbrev=0 --tags --always)
Target     = $(shell basename $(abspath $(dir $$PWD)))
Suffix     =
GOFILES    = $(shell find . -name "*.go" -type f ! -path "./vendor/*" ! -path "*/bindata.go")

ifeq ($(OS),Windows_NT)
	OSName = windows
	Suffix = .exe
else
	OSName = $(shell echo $(shell uname -s) | tr '[:upper:]' '[:lower:]')
endif

.PHONY: ${OSName}
${OSName}: bindata
	GOOS=$@ go build -v -o release/${Target}-$@${Suffix} -ldflags \
	"-s -w -X main.BuildStamp=${BuildStamp} -X main.GitHash=${GitHash} -X main.Version=${Version}"

.PHONY: dev
dev:
	GOOS=${OSName} go build -v -o release/${Target}-${OSName}${Suffix} -tags "$@" -ldflags \
	"-s -w -X main.BuildStamp=${BuildStamp} -X main.GitHash=${GitHash} -X main.Version=${Version}"

.PHONY: release
release: clean
	xgo -v -out ${Target}-${Version} --targets=windows/*,darwin/*,linux/* -ldflags "-s -w -X main.BuildStamp=${BuildStamp} -X main.GitHash=${GitHash} -X main.Version=${Version}" github.com/tosone/panx
	mkdir release
	mv ${Target}-${Version}-* release

.PHONY: bindata
bindata:
	cd public && go-bindata -o=../bindata/public/bindata.go -pkg=public ./...
	cd templates && go-bindata -o=../bindata/templates/bindata.go -pkg=templates ./...

.PHONY: authors
authors:
	printf "Authors\n=======\n\nProject's contributors:\n\n" > AUTHORS.md
	git log --raw | grep "^Author: " | cut -d ' ' -f2- | cut -d '<' -f1 | sed 's/^/- /' | sort | uniq >> AUTHORS.md

.PHONY: lint
lint:
	gometalinter.v2 --deadline 5m ./...

.PHONY: misspell-check
misspell-check:
	@hash misspell > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		go get -u github.com/client9/misspell/cmd/misspell; \
	fi
	misspell -error -i unknwon $(GOFILES)

.PHONY: clean
clean:
	$(RM) -r release *.db *.db-journal
