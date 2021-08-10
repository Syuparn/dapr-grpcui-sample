FROM golang:1.16

RUN apt-get update && \
    apt-get install -y curl unzip

WORKDIR /proto

ENV PROTOC_VER=3.17.3
ENV PROTOC_ZIP=protoc-${PROTOC_VER}-linux-x86_64.zip
RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VER}/${PROTOC_ZIP} && \
    unzip -o ${PROTOC_ZIP} -d . && \
    mv ./bin/protoc /usr/bin/protoc && \
    mv ./include /usr/bin

WORKDIR /evans

ENV EVANS_VER=0.9.3
RUN mkdir evans && \
    curl -sL https://github.com/ktr0731/evans/releases/download/${EVANS_VER}/evans_linux_amd64.tar.gz \
    | tar zx -C evans && \
    mv evans/evans /usr/bin/evans

WORKDIR /go/src

COPY . .

#RUN go install github.com/golang/protobuf/protoc-gen-go

RUN go mod download

RUN go build
