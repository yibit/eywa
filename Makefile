NAME := eywa
VERSION := $(shell cat ./VERSION |awk 'NR==1 { print $1; }')
GOMODULES = ./...
MYHOME := $(PWD)
GOFILES = $(shell cd $(NAME) && go list $(GOMODULES) |grep -v /vendor/)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
COMMIT = $(shell git rev-parse --short HEAD)
RELEASEDATE=$(shell date '+%Y%m%d%H%M%S')
LDFLAGS = "-X main.Version=$(NAME)-$(VERSION)-$(BRANCH)-$(COMMIT)-$(RELEASEDATE)-$@ -s -w"

all: usage

usage:
	@echo "Usage:                                              "
	@echo "                                                    "
	@echo "    make command                                    "
	@echo "                                                    "
	@echo "The commands are:                                   "
	@echo "                                                    "
	@echo "    build       compile packages and dependencies   "
	@echo "    dev         run go build -mod=vendor            "
	@echo "    debug       run go build -tags debug            "
	@echo "    test        run go test                         "
	@echo "    clean       remove object files                 "
	@echo "    fmt         run gofmt on package sources        "
	@echo "    cov         run go tool cover                   "
	@echo "    release     release a version                   "
	@echo "                                                    "

setup: clean vendor build

build:
	@sh tools/git-status-check.sh
	cd $(NAME) && CGO_ENABLED=0 go build -ldflags=$(LDFLAGS) -v -o $(MYHOME)/bin/$(NAME)

rebuild:
	@sh tools/git-status-check.sh
	cd $(NAME) && CGO_ENABLED=0 go build -ldflags=$(LDFLAGS) -v -a -o $(MYHOME)/bin/$(NAME)

ox:
	cd $(NAME) && CGO_ENABLED=0 garble build -ldflags=$(LDFLAGS) -v -a -o $(MYHOME)/bin/$(NAME)

debug:
	cd $(NAME) && CGO_ENABLED=0 go build -tags debug -ldflags=$(LDFLAGS) -v -x -a -o $(MYHOME)/bin/$(NAME) -gcflags "all=-N -l"

dev: vendor
	cd $(NAME) && go build -mod=vendor -ldflags=$(LDFLAGS) -o $(MYHOME)/bin/$(NAME)

vendor:
	cd $(NAME) && go mod vendor

mod:
	cd $(NAME) && go mod download

tidy:
	cd $(NAME) && go mod tidy && go mod vendor

cov:
	cd $(NAME) && go test -v $(GOMODULES) -coverprofile=coverage.out
	cd $(NAME) && go tool cover -html=coverage.out -o coverage.html

fmt:
	cd $(NAME) && go fmt $(GOFILES)

lint:
	cd $(NAME) && golint $(GOFILES)

# https://github.com/jorisroovers/gitlint
gitlint:
	gitlint

mdlint:
	markdownlint -c .markdownlint.yml book/src

# https://github.com/golangci/golangci-lint
golangci-lint:
	cd $(NAME) && golangci-lint run -v

# https://www.praetorian.com/blog/introducing-gokart/
gokart:
	gokart scan $(NAME)

# https://github.com/mvdan/sh
shfmt:
	@echo ">> formatting shell scripts"
	@shfmt -i 4 -ci -w -s $(shell find . -type f -name "*.sh" -not -path "*vendor*")

# brew install shellcheck
# https://github.com/koalaman/shellcheck
shellcheck:
	shellcheck $(shell find . -type f -name "*.sh" -not -path "*vendor*")

# https://github.com/client9/misspell
misspell:
	@misspell -error $(shell find $(NAME) -type f -name "*.go" |grep -v vendor)

test:
	cd $(NAME) && go test -v -failfast -race -coverpkg=./... -covermode=atomic -coverprofile=coverage.out $(GOFILES) -run . -timeout=2m

#delve: go get github.com/go-delve/delve/cmd/dlv
dlv:
	$(GOPATH)/bin/dlv exec bin/$(NAME)

# https://github.com/wagoodman/dive
dive:
	$(GOPATH)/bin/dive $(NAME)-$(VERSION)

tag:
	git tag -a $(VERSION) -m "Release: $(VERSION)" || true
	git push origin $(VERSION)

# https://goreleaser.com/
release:
	cd $(NAME) && goreleaser release --rm-dist

.PHONY: clean check distclean doc fmt test release

clean:
	rm -f $(NAME).svg $(NAME).png $(NAME)/coverage.* trace.out dive.log *.tags
	find . -name \*~ -type f |xargs -I {} rm -f {}
	find . -type f |grep -E "\._.*" |xargs -I {} rm -f {}

distclean: clean
	rm -f bin/$(NAME)
