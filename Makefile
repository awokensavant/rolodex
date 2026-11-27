VERSION=`cat VERSION`
REVISION=`git rev-parse --short HEAD`
BUILDTIME=`date +%FT%T%z`
API=`cat API_VERSION`

BUILD_DIR=.build
BINARY=rolodex
CONFIG_NAME=config-rolodex

LDFLAGS=-ldflags "-X github.com/awokensavant/rolodex.Version=${VERSION} -X github.com/awokensavant/rolodex.Revision=${REVISION} -X github.com/awokensavant/rolodex.BuildTime=${BUILDTIME} -X github.com/awokensavant/rolodex.API=${API}"

.PHONY: build
# build: builds the rolodex binary
build: clean
	@echo "Building rolodex for current architecture"
	@go build ${LDFLAGS} -o ${BUILD_DIR}/${BINARY} cmd/rolodex.go
	@cp workdir/config.yaml ${BUILD_DIR}/${CONFIG_NAME}.yaml
	@echo "Done building ${BINARY} @ ${BUILD_DIR}"

.PHONY: clean
# clean: cleans the .build directory
clean:
	@echo "Cleaning ${BUILD_DIR}"
	@rm -rf ${BUILD_DIR}
	@echo "Done cleaning ${BUILD_DIR}"

.PHONY: help
# help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^#/ /p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'