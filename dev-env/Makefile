GOLANG=build-onamaeddns-go
DC=docker-compose

CURRENT_UID := $(shell id -u)
CURRENT_GID := $(shell id -g)
BUILD_VERSION := $(shell git describe --tags)

export CURRENT_UID
export CURRENT_GID
export BUILD_VERSION

SRCS := $(shell find . -type f)

.PHONY: all
all: run gomod gobuild stop## test & build

.PHONY: build
build: $(SRCS) ## build to all container
	$(DC) build

.PHONY: run
run: $(SRCS) ## start up to all container
	$(DC) up -d

.PHONY: stop
stop: $(SRCS) ## down to to all container
	$(DC) stop

.PHONY: test
test:
	$(DC) exec -T $(GOLANG) pwd

.PHONY: clean
clean: ## stop container and cleanup data
	$(DC) down --rmi all

.PHONY: go
go: ## run golang
	$(DC) exec -T $(GOLANG) make
.PHONY: gobuild
gobuild: ## build golang
	$(DC) exec -T $(GOLANG) make build
.PHONY: gotest
gotest: ## test golang
	$(DC) exec -T $(GOLANG) make test
.PHONY: gomod
gomod: ## run golang package manager
	$(DC) exec -T $(GOLANG) make mod
.PHONY: gomodinit
gomodinit: ## init golang package manager
	$(DC) exec -T $(GOLANG) make modinit
.PHONY: goclean
goclean: ## clean golang
	$(DC) exec -T $(GOLANG) make clean

.PHONY: help
	all: help
help: ## help
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {\
		printf "\033[36m%-30s\033[0m %s\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)
