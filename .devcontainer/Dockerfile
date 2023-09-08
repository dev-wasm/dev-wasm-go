
FROM --platform=$BUILDPLATFORM ubuntu:jammy

ARG TARGETPLATFORM
ARG BUILDPLATFORM

ENV TINY_GO_VERSION=0.29.0
ENV GOLANG_VERSION=1.21.1

RUN apt-get update && apt-get install -yy -q curl

RUN archi=$(echo $TARGETPLATFORM | cut -d / -f2) && \
    curl https://github.com/tinygo-org/tinygo/releases/download/v${TINY_GO_VERSION}/tinygo_${TINY_GO_VERSION}_$archi.deb -L --output tinygo_${TINY_GO_VERSION}_$archi.deb && \ 
    dpkg -i tinygo_${TINY_GO_VERSION}_$archi.deb && \ 
    rm tinygo_${TINY_GO_VERSION}_$archi.deb && \ 
    curl --output go${GOLANG_VERSION}.linux-$archi.tar.gz https://go.dev/dl/go${GOLANG_VERSION}.linux-$archi.tar.gz -L && \ 
    tar -C /usr/local -xzf go${GOLANG_VERSION}.linux-$archi.tar.gz && \ 
    rm go${GOLANG_VERSION}.linux-$archi.tar.gz

