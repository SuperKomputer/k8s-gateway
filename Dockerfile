FROM alpine:latest

ADD gateway /gateway

ENTRYPOINT /gateway
