FROM --platform=linux/amd64 ubuntu:20.04 AS builder

ENV VERSION 1.0.0

ENV LANG en_US.UTF-8
ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && \
    apt-get install -y curl tzdata && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* && \
    ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

ADD ./bin/Linux_x86_64/onamaeddns /usr/local/bin/onamaeddns
ADD ./docker-in/exec_ddns.sh /usr/local/bin/exec_ddns.sh

RUN chmod +x /usr/local/bin/onamaeddns && \
    chmod +x /usr/local/bin/exec_ddns.sh

ENTRYPOINT ["/usr/local/bin/exec_ddns.sh"]
