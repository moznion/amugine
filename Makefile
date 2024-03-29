PACKAGE=github.com/moznion/amugine
PKGS := $(shell go list ./...)
GOVERSION=$(shell go version)
GOOS=$(word 1,$(subst /, ,$(lastword $(GOVERSION))))
GOARCH=$(word 2,$(subst /, ,$(lastword $(GOVERSION))))
RELEASE_DIR=bin
REVISION=$(shell git rev-parse --verify HEAD)

build: $(RELEASE_DIR)/amugine_$(GOOS)_$(GOARCH)

all: check clean build-linux-amd64 build-linux-386 build-darwin-amd64 build-darwin-386 build-windows-amd64 build-windows-386

build-linux-amd64:
	@$(MAKE) build GOOS=linux GOARCH=amd64

build-linux-386:
	@$(MAKE) build GOOS=linux GOARCH=386

build-darwin-amd64:
	@$(MAKE) build GOOS=darwin GOARCH=amd64

build-darwin-386:
	@$(MAKE) build GOOS=darwin GOARCH=386

build-windows-amd64:
	@$(MAKE) build GOOS=windows GOARCH=amd64

build-windows-386:
	@$(MAKE) build GOOS=windows GOARCH=386

$(RELEASE_DIR)/amugine_$(GOOS)_$(GOARCH):
ifndef VERSION
	@echo '[ERROR] $$VERSION must be specified'
	exit 255
endif
	go build -ldflags "-X $(PACKAGE).rev=$(REVISION) -X $(PACKAGE).ver=$(VERSION)" \
		-o $(RELEASE_DIR)/amugine_$(GOOS)_$(GOARCH)_$(VERSION) cmd/amugine/amugine.go

clean:
	rm -rf $(RELEASE_DIR)/amugine_*

check: test lint vet

test:
	go test $(PKGS)

lint:
	golint -set_exit_status $(PKGS)

vet:
	go vet $(PKGS)

