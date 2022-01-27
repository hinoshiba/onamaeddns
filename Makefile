
SRCS := $(shell find . -name '*.go' -type f)

.PHONY: all
all: build-bin build-docker ## build binary and build docker.
	docker build -t onamaeddns:debug .
	make builded_flg.unlock

.PHONY: clean
clean: builded_flg.unlock ## cleanup
	docker rmi hinoshiba/onamaeddns:debug
	make -C dev-env clean
	rm -rf ./vendor/ || exit 0
	rm -rf ./bin/ || exit 0

.PHONY: build-docker
build-docker: Dockerfile docker-in/exec_ddns.sh builded_flg.mtx ## build docker image.
	docker-compose build

.PHONY: build-bin
build-bin: builded_flg.mtx ## build binary.

builded_flg.mtx: $(SRCS)
	touch builded_flg.mtx
	# build bin/Linux_x86_64/onamaeddns bin/Darwin_aarch64/onamaeddns
	make -C dev-env

.PHONY: builded_flg.unlock
builded_flg.unlock:
	rm builded_flg.mtx || exit 0

.PHONY: help
help: ## help
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {\
		printf "\033[36m%-30s\033[0m %s\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)
