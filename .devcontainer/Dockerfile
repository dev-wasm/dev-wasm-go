
FROM --platform=$BUILDPLATFORM ubuntu:jammy

ARG TARGETPLATFORM
ARG BUILDPLATFORM

ENV TINY_GO_VERSION=0.30.0
ENV GOLANG_VERSION=1.21.1
ENV WASM_TOOLS_VERSION=1.0.54

RUN apt-get update && apt-get install -yy -q curl

RUN archi=$(echo $TARGETPLATFORM | cut -d / -f2) && \
    curl https://github.com/tinygo-org/tinygo/releases/download/v${TINY_GO_VERSION}/tinygo_${TINY_GO_VERSION}_$archi.deb -L --output tinygo_${TINY_GO_VERSION}_$archi.deb && \ 
    dpkg -i tinygo_${TINY_GO_VERSION}_$archi.deb && \ 
    rm tinygo_${TINY_GO_VERSION}_$archi.deb && \ 
    curl --output go${GOLANG_VERSION}.linux-$archi.tar.gz https://go.dev/dl/go${GOLANG_VERSION}.linux-$archi.tar.gz -L && \ 
    tar -C /usr/local -xzf go${GOLANG_VERSION}.linux-$archi.tar.gz && \ 
    rm go${GOLANG_VERSION}.linux-$archi.tar.gz

RUN curl --output wasm-tools-${WASM_TOOLS_VERSION}-x86_64-linux.tar.gz -L https://github.com/bytecodealliance/wasm-tools/releases/download/wasm-tools-${WASM_TOOLS_VERSION}/wasm-tools-${WASM_TOOLS_VERSION}-x86_64-linux.tar.gz
RUN tar xzf wasm-tools-${WASM_TOOLS_VERSION}-x86_64-linux.tar.gz
RUN mv wasm-tools-${WASM_TOOLS_VERSION}-x86_64-linux/wasm-tools /usr/local/bin/wasm-tools
RUN apt install -y -q make

