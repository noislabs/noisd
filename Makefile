.PHONY: build proto check_go_version install
#!/usr/bin/make -f

PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true
SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\@/g')
TM_VERSION := $(shell go list -m github.com/tendermint/tendermint | sed 's:.* ::')
BUILDDIR ?= $(CURDIR)/build
DOCKER := $(shell which docker)

export GO111MODULE = on

GO_MAJOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f1)
GO_MINOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f2)

# process build tags
build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq (cleveldb,$(findstring cleveldb,$(NOIS_BUILD_OPTIONS)))
  build_tags += gcc
else ifeq (rocksdb,$(findstring rocksdb,$(NOIS_BUILD_OPTIONS)))
  build_tags += gcc
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=nois \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=noisd \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)" \
		  -X github.com/tendermint/tendermint/version.TMCoreSemVer=$(TM_VERSION)

ifeq (cleveldb,$(findstring cleveldb,$(NOIS_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
else ifeq (rocksdb,$(findstring rocksdb,$(NOIS_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=rocksdb
endif
ifeq (,$(findstring nostrip,$(NOIS_BUILD_OPTIONS)))
  ldflags += -w -s
endif
ifeq ($(LINK_STATICALLY),true)
	ldflags += -linkmode=external -extldflags "-Wl,-z,muldefs -static"
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
# check for nostrip option
ifeq (,$(findstring nostrip,$(NOIS_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif

check_go_version:
ifeq ($(GO_MINOR_VERSION),19)
	@echo "Go version: $(GO_MAJOR_VERSION).$(GO_MINOR_VERSION). Ok."
else ifeq ($(GO_MINOR_VERSION),20)
	@echo "Go version: $(GO_MAJOR_VERSION).$(GO_MINOR_VERSION). Ok."
else ifeq ($(GO_MINOR_VERSION),21)
	@echo "Go version: $(GO_MAJOR_VERSION).$(GO_MINOR_VERSION). Ok."
else
	@echo "ERROR: Go version 1.19, 1.20 or 1.21 is required for this version of Nois"
	exit 1
endif

all: install

install: check_go_version
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/noisd

build: check_go_version
	go build $(BUILD_FLAGS) -o build/noisd ./cmd/noisd

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	GO111MODULE=on go mod verify

# look into .golangci.yml for enabling / disabling linters
lint:
	@echo "--> Running linter"
	@golangci-lint run --tests=false --skip-dirs="simapp"
	@go mod verify


full-lint: lint
	@gosec -exclude-dir=cmd ./...

build-linux:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build $(BUILD_FLAGS) -o build/noisd github.com/noislabs/noisd/cmd/noisd

test:
	go test -v -race github.com/noislabs/noisd/x/...

.PHONY: test build-linux docker-test lint build install format

format:
	gofumpt -l -w .
###############################################################################
###                                Protobuf                                 ###
###############################################################################


BUF_IMAGE=bufbuild/buf@sha256:3cb1f8a4b48bd5ad8f09168f10f607ddc318af202f5c057d52a45216793d85e5 #v1.4.0
DOCKER_BUF := $(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(BUF_IMAGE)


###############################################################################
###                                Protobuf                                 ###
###############################################################################
PROTO_BUILDER_IMAGE=tendermintdev/sdk-proto-gen:v0.2
PROTO_FORMATTER_IMAGE=tendermintdev/docker-build-proto@sha256:aabcfe2fc19c31c0f198d4cd26393f5e5ca9502d7ea3feafbfe972448fee7cae

proto-all: proto-format proto-lint proto-gen format

proto-gen:
	@echo "Generating Protobuf files"
	$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(PROTO_BUILDER_IMAGE) sh ./scripts/protocgen.sh

proto-format:
	@echo "Formatting Protobuf files"
	$(DOCKER) run --rm -v $(CURDIR):/workspace \
	--workdir /workspace $(PROTO_FORMATTER_IMAGE) \
	find ./ -name *.proto -exec clang-format -i {} \;

proto-swagger-gen:
	@./scripts/protoc-swagger-gen.sh

# Show the version that `make build` is using
.PHONY: version
version:
	@echo "$(VERSION)"
