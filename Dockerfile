# Cross-build from Intel or ARM to x86_64. The node never uses an ARM build due to https://github.com/CosmWasm/cosmwasm/issues/1628
#   Build:    docker buildx build --platform linux/amd64 --pull --tag "noislabs/noisd:$(make version)" . --load
#   Publish:  docker buildx build --platform linux/amd64 --pull --tag "noislabs/noisd:$(make version)" . --push
#
# Run
#   show version:       docker run --rm noislabs/noisd:manual
#   libwasmvm version:  docker run --rm noislabs/noisd:manual noisd query wasm libwasmvm-version
#   shell:              docker run --rm -it noislabs/noisd:manual sh
FROM golang:1.20.2-alpine3.17 AS go-builder

# this comes from standard alpine nightly file
#  https://github.com/rust-lang/docker-rust-nightly/blob/master/alpine3.12/Dockerfile
# with some changes to support our toolchain, etc
RUN set -eux; apk add --no-cache ca-certificates build-base;

RUN apk add git
# NOTE: add these to run with LEDGER_ENABLED=true
# RUN apk add libusb-dev linux-headers

# See https://github.com/CosmWasm/wasmvm/releases
# Copy the library you want to the final location that will be found by the linker flag `-lwasmvm_muslc`
RUN APK_ARCH="$(apk --print-arch)"; \
  echo "Detected architecture: $APK_ARCH"; \
  case "$APK_ARCH" in \
    aarch64) \
      wget -O /lib/libwasmvm_muslc.a https://github.com/CosmWasm/wasmvm/releases/download/v1.2.5/libwasmvm_muslc.aarch64.a \
        && sha256sum /lib/libwasmvm_muslc.a | grep 741c5862c237c17edd32dcebdd1d43d40d1684afe049b0f0aba9da8f556a6285 ;; \
    x86_64) \
      wget -O /lib/libwasmvm_muslc.a https://github.com/CosmWasm/wasmvm/releases/download/v1.2.5/libwasmvm_muslc.x86_64.a \
        && sha256sum /lib/libwasmvm_muslc.a | grep cfbf72e719a33328fbb196d88795f1f2a0a9b2288d4c882b5467fba39f8de86d ;; \
  esac;

WORKDIR /code
COPY . /code/

# force it to use static lib (from above) not standard libgo_cosmwasm.so file
RUN LEDGER_ENABLED=false BUILD_TAGS=muslc LINK_STATICALLY=true make build
RUN echo "Ensuring binary is statically linked ..." \
  && (file /code/build/noisd | grep "statically linked")

# --------------------------------------------------------
FROM alpine:3.17 as noisd

COPY --from=go-builder /code/build/noisd /usr/bin/noisd

# jq is required in the setup script at runtime
RUN apk add jq

COPY docker/scripts/* /usr/local/bin/
RUN chmod +x /usr/local/bin/*.sh

WORKDIR /opt

# tendermint p2p
EXPOSE 26656
# tendermint rpc
EXPOSE 26657

CMD ["/usr/bin/noisd", "version"]
