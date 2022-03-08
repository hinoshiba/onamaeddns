FROM golang:1.17.8 AS builder

#ENV http_proxy http://proxy.example.com:8080/
#ENV https_proxy http://proxy.example.com:8080/
ENV DEBIAN_FRONTEND=noninteractive
ENV DEBCONF_NOWARNINGS=yes

RUN apt-get update && \
    apt-get install -y wget make gcc && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN ln -s /opt/go/bin/go /usr/bin/ && \
    mkdir -p /go/ && \
    mkdir /.cache && \
    chmod 777 /.cache && \
    mkdir -p /go/bin && \
    mkdir -p /go/pkg && \
    chmod 777 -R /go/bin && \
    chmod 777 -R /go/pkg
RUN ln -s /go/src/Makefile.4builder /go/Makefile

WORKDIR /go
