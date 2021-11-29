GOLANG=build-onamaeddns-go
DC=docker-compose

CURRENT_UID := $(shell id -u)
CURRENT_GID := $(shell id -g)
BUILD_VERSION := $(shell git describe --tags)

export CURRENT_UID
export CURRENT_GID
export BUILD_VERSION

SRCS := $(shell find . -type f)

all: build run gomod gobuild stop## test & build

build: $(SRCS) ## build to all container
	$(DC) build

run: $(SRCS) ## start up to all container
	$(DC) up -d

stop: $(SRCS) ## down to to all container
	$(DC) down

.PHONY: clean
clean: ## stop container and cleanup data
	$(DC) down --rmi all

go: ## run golang
	$(DC) exec $(GOLANG) make
gobuild: ## build golang
	$(DC) exec $(GOLANG) make build
goxbuild: ## cross build golang
	$(DC) exec $(GOLANG) make build
	$(DC) exec $(GOLANG) make build-mac
gotest: ## test golang
	$(DC) exec $(GOLANG) make test
gomod: ## run golang package manager
	$(DC) exec $(GOLANG) make mod
gomodinit: ## init golang package manager
	$(DC) exec $(GOLANG) make modinit
goclean: ## clean golang
	$(DC) exec $(GOLANG) make clean

.PHONY: help
	all: help
help: ## help
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {\
		printf "\033[36m%-30s\033[0m %s\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)
