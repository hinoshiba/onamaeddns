
SRCS := $(shell find . -name '*.go' -type f)

.PHONY: all
all: build-bin build-docker ## build binary and build docker.
	docker-compose build

.PHONY: clean
clean: ## cleanup
	docker-compose down
	make -C dev-env clean
	rm -rf ./vendor/
	rm -rf ./bin/

.PHONY: build-docker
build-docker: Dockerfile docker-in/exec_ddns.sh bin/Linux_x86_64/onamaeddns ## build docker image.
	docker-compose build

.PHONY: build-bin
build-bin: bin/Linux_x86_64/onamaeddns bin/Darwin_aarch64/onamaeddns ## build binary.


bin/Linux_x86_64/onamaeddns: $(SRCS)
	make -C dev-env
bin/Darwin_aarch64/onamaeddns: $(SRCS)
	make -C dev-env

help: ## help
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {\
		printf "\033[36m%-30s\033[0m %s\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)
