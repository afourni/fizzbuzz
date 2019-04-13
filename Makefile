PROJECT_NAME := "fizzbuzz"

# Environment variable to enable go module
export GO111MODULE=on

# Environment variables required to build the docker image statically
# in order to use Docker scratch image
export CGO_ENABLED=0
export GOOS=linux

PKG := "github.com/afourni/$(PROJECT_NAME)"
PKG_PATH := $(GOPATH)/src/$(PKG)
PKG_LIST := $(shell go list ${PKG_PATH}/... | grep -v /vendor/)

all: test bench build docker

# Run tests
test:
	@echo '*** Testing' $(PKG)
	@go test -v -cover ${PKG_LIST}

# Run benchmarks
bench:
	@echo '*** Benchmarking' $(PKG)
	@go test -bench=. ${PKG_LIST}

# Build the project
build:
	@echo '*** Building' $(PKG)
	@go build -a -ldflags '-extldflags "-static"' -i -v $(PKG_PATH)

# Build the Docker image
docker:
	@echo '*** Building the Docker image' $(PKG)
	@docker build -t $(PKG):latest $(PKG_PATH)

# Remove the previous build and the Docker image
clean:
	@echo '*** Cleaning' $(PKG)
	@rm -f $(PROJECT_NAME)
	@docker rmi -f $(PKG):latest