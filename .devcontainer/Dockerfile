FROM brendanburns/wasm-dev-base:0.0.7

ENV TINY_GO_VERSION=0.26.0
ENV GOLANG_VERSION=1.19.3

RUN curl https://github.com/tinygo-org/tinygo/releases/download/v${TINY_GO_VERSION}/tinygo_${TINY_GO_VERSION}_amd64.deb -L --output tinygo_${TINY_GO_VERSION}_amd64.deb && \
    dpkg -i tinygo_${TINY_GO_VERSION}_amd64.deb && \
    rm tinygo_${TINY_GO_VERSION}_amd64.deb && \
    curl --output go${GOLANG_VERSION}.linux-amd64.tar.gz https://go.dev/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz -L && \
    tar -C /usr/local -xzf go${GOLANG_VERSION}.linux-amd64.tar.gz && \
    rm go${GOLANG_VERSION}.linux-amd64.tar.gz

