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

test-unit: 
	@echo "Executing unit tests..."
	@go test -mod=readonly -v -coverprofile coverage.txt ./...
.PHONY: test-unit