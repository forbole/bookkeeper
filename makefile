VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT  := $(shell git log -1 --format='%H')


###############################################################################
###                                Build flags                              ###
###############################################################################

LD_FLAGS = -X github.com/forbole/bookkeeper/version.Version=$(VERSION) \
	-X github.com/forbole/bookkeeper/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(LD_FLAGS)'


###############################################################################
###                                  Build                                  ###
###############################################################################

build: go.sum
ifeq ($(OS),Windows_NT)
	@echo "building bookkeeper binary..."
	@go build -mod=readonly $(BUILD_FLAGS) -o build/bdbookkeeper.exe ./cmd/bdbookkeeper
else
	@echo "building bookkeeper binary..."
	@go build -mod=readonly $(BUILD_FLAGS) -o build/bdbookkeeper ./cmd/bdbookkeeper
endif
.PHONY: build


###############################################################################
###                                 Install                                 ###
###############################################################################

install: go.sum
	@echo "installing bookkeeper binary..."
	@go install -mod=readonly $(BUILD_FLAGS) ./cmd/bookkeeper
.PHONY: install

###############################################################################
###                                 Test                                 ###
###############################################################################
stop-docker-test:
	@echo "Stopping Docker container..."
	@docker stop bdjuno-test-db || true && docker rm bdjuno-test-db || true
.PHONY: stop-docker-test

start-docker-test: stop-docker-test
	@echo "Starting Docker container..."
	@docker run --name bdjuno-test-db -e POSTGRES_USER=bdjuno -e POSTGRES_PASSWORD=password -e POSTGRES_DB=bdjuno -d -p 5433:5432 postgres 
.PHONY: start-docker-test

test-unit: start-docker-test
	@echo "Executing unit tests..."
	@go test -mod=readonly -v -coverprofile coverage.txt ./... 
.PHONY: test-unit


lint:
	golangci-lint run --out-format=tab

lint-fix:
	golangci-lint run --fix --out-format=tab --issues-exit-code=0
.PHONY: lint lint-fix