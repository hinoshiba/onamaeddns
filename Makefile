NAME=onamaeddns
PRJ=src/$(NAME)

#GOENV=GOPATH=$(CURDIR)
GOCMD=go
GOBUILD=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod

BUILD_FLGS= -tags netgo -installsuffix netgo -ldflags='-extldflags="static"'
WIN_OPT=GO111MODULE=on CGO_ENABLED=1 GOOS=windows GOARCH=amd64
MAC_OPT=GO111MODULE=on GOOS=darwin GOARCH=arm64

SRCS := $(shell find . -name '*.go' -type f)
BINS := $(shell test -d ./bin && find ./bin/ -type f)

all: test build ## test & build

build: $(SRCS) ## build to linux binary
	$(GOBUILD) $(BUILD_FLGS) exec/...
	#cd $(CURDIR)/$(PRJ); $(MAC_OPT) $(GOBUILD) $(BUILD_FLGS) exec/...
	$(MAC_OPT) $(GOBUILD) $(BUILD_FLGS) exec/...

.PHONY: test
test: ## run test
	$(GOTEST) -count=1 ./$(PRJ)/...

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
