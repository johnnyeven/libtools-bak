FROM golang:1.10-alpine

RUN sed -i "s|http://dl-cdn.alpinelinux.org|http://mirrors.aliyun.com|g" /etc/apk/repositories

RUN apk add --no-cache --virtual .build-deps \
	    curl bash vim

ENV GODEBUG=netdns=cgo

WORKDIR /etc/service