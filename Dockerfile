FROM golang:alpine as builder

WORKDIR /go/src/server
COPY . .

RUN echo "https://mirrors.aliyun.com/alpine/v3.6/main/" > /etc/apk/repositories && apk add gcc g++ make libffi-dev openssl-dev libtool git
RUN make setup && make build

FROM alpine:latest
LABEL MAINTAINER="lirui@thooh.com"

WORKDIR /go/src/server

COPY --from=builder /go/src/server/build ./

EXPOSE 8080
#VOLUME /data/conf
COPY configs/config.example.toml /data/conf/config.toml

CMD ["./server", "-conf", "/data/conf"]
