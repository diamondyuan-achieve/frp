FROM ubuntu:16.04 as frp
MAINTAINER DiamondYuan <541832074>

ARG FRP_VERSION=0.13.0

RUN apt update \
    && apt install -y wget

WORKDIR /tmp
RUN set -x \
    && wget https://github.com/fatedier/frp/releases/download/v${FRP_VERSION}/frp_${FRP_VERSION}_linux_amd64.tar.gz \
    && tar -zxf frp_${FRP_VERSION}_linux_amd64.tar.gz \
    && mv frp_${FRP_VERSION}_linux_amd64 /var/frp 

FROM alpine:3.6
COPY --from=frp /var/frp/ /frp/

ENTRYPOINT /frp/frps -c /frp/frps.ini