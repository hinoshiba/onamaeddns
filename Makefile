GOENV=GOPATH=$(CURDIR)
GOCMD=go
GOBUILD=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod

BUILD_FLGS= -tags netgo -installsuffix netgo -ldflags='-extldflags="static"'
LINUX_OPT= GOOS=linux GOARCH=amd64
MAC_OPT= GOOS=darwin GOARCH=arm64

SRCS := $(shell find . -name '*.go' -type f)
BINS := $(shell test -d ./bin && find ./bin/ -type f)

all: test build ## test & build

build: $(SRCS) ## build to linux binary
	$(LINUX_OPT) $(GOBUILD) $(BUILD_FLGS) ./exec/onamaeddns/onamaeddns.go
	mkdir -p /go/src/bin/Linux_x86_64/
	mv /go/src/bin/onamaeddns /go/src/bin/Linux_x86_64/
	$(MAC_OPT) $(GOBUILD) $(BUILD_FLGS) ./exec/onamaeddns/onamaeddns.go
	mkdir -p /go/src/bin/Darwin_aarch64/
	mv /go/src/bin/onamaeddns /go/src/bin/Darwin_aarch64/

.PHONY: test
test: ## run test
	$(GOTEST) -count=1 ./onamaeddns_test.go

.PHONY: clean
clean: $(BINS) ## cleanup
	$(GOCLEAN)
	rm -f $(BINS)

mod: go.mod ## mod ensure
	$(GOMOD) tidy
	$(GOMOD) vendor
modinit: ## mod init
	$(GOMOD) init

.PHONY: help
	all: help
help: ## help
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {\
		printf "\033[36m%-30s\033[0m %s\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)
