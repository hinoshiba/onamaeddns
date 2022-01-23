FROM --platform=linux/amd64 ubuntu:20.04 AS builder

ENV GOVERSION 1.17.6

ENV LANG en_US.UTF-8
#ENV http_proxy http://proxy.example.com:8080/
#ENV https_proxy http://proxy.example.com:8080/
ENV DEBIAN_FRONTEND=noninteractive

RUN apt update && \
    apt install -y git curl wget tzdata make gcc && \
    apt clean && \
    rm -rf /var/lib/apt/lists/* && \
    ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

RUN cd /opt && wget -nv --show-progress --progress=bar:force:noscroll https://golang.org/dl/go${GOVERSION}.linux-amd64.tar.gz && \
    tar zxf go${GOVERSION}.linux-amd64.tar.gz && rm go${GOVERSION}.linux-amd64.tar.gz && \
    ln -s /opt/go/bin/go /usr/bin/ && \
    mkdir -p /go/ && \
    mkdir /.cache && \
    chmod 777 /.cache && \
    mkdir -p /go/pkg && \
    chmod 777 -R /go/pkg
RUN ln -s /go/src/Makefile.inbuilder /go/Makefile

WORKDIR /go
