# Build image
#   on Intel: docker build --target noisd -t noislabs/noisd:manual .
#   on ARM:   docker build --target noisd -t noislabs/noisd:manual --build-arg arch=aarch64 .
#
# Run
#   show version:       docker run --rm noislabs/noisd:manual
#   libwasmvm version:  docker run --rm noislabs/noisd:manual /usr/bin/noisd query wasm libwasmvm-version
#   shell:              docker run --rm -it noislabs/noisd:manual /bin/sh
FROM golang:1.20.1-alpine3.17 AS go-builder
ARG arch=x86_64

# this comes from standard alpine nightly file
#  https://github.com/rust-lang/docker-rust-nightly/blob/master/alpine3.12/Dockerfile
# with some changes to support our toolchain, etc
RUN set -eux; apk add --no-cache ca-certificates build-base;

RUN apk add git
# NOTE: add these to run with LEDGER_ENABLED=true
# RUN apk add libusb-dev linux-headers

WORKDIR /code
COPY . /code/
# See https://github.com/CosmWasm/wasmvm/releases
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.2.0/libwasmvm_muslc.aarch64.a /lib/libwasmvm_muslc.aarch64.a
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.2.0/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.x86_64.a
RUN sha256sum /lib/libwasmvm_muslc.aarch64.a | grep cba4b334893456c64df177939cbdd09afe4812432c02ae37d60d69a111b1b50d
RUN sha256sum /lib/libwasmvm_muslc.x86_64.a | grep 6f87082f7a62602f9725d529677f330b9c4dd4607887be52a86328c6c919495b

# Copy the library you want to the final location that will be found by the linker flag `-lwasmvm_muslc`
RUN cp /lib/libwasmvm_muslc.${arch}.a /lib/libwasmvm_muslc.a

# force it to use static lib (from above) not standard libgo_cosmwasm.so file
RUN LEDGER_ENABLED=false BUILD_TAGS=muslc LINK_STATICALLY=true make build
RUN echo "Ensuring binary is statically linked ..." \
  && (file /code/build/noisd | grep "statically linked")

# --------------------------------------------------------
FROM alpine:3.17 as noisd

COPY --from=go-builder /code/build/noisd /usr/bin/noisd

# COPY docker/* /opt/
# RUN chmod +x /opt/*.sh

WORKDIR /opt

# rest server
EXPOSE 1317
# tendermint p2p
EXPOSE 26656
# tendermint rpc
EXPOSE 26657

CMD ["/usr/bin/noisd", "version"]

FROM noisd as noisd-ci

# USER root

RUN apk add jq

WORKDIR /opt

COPY docker/scripts-ci/* /opt
RUN chmod +x /opt/*.sh

CMD [ "/opt/run.sh" ]
