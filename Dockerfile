ARG GO_VERSION="1.18"
ARG ALPINE_VERSION="3.16"


#
## Dependencies Stage
FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS dependencies-stage
# ENV GOPROXY="http://192.168.1.182:90/"
ENV GO111MODULE="on"
WORKDIR /go/src/royfuwei.com/tools/go-clean-arch

RUN apk update && apk --no-cache upgrade musl && apk add --no-cache tzdata git

COPY go.mod .
COPY go.sum .

RUN go mod download 


#
## Build Stage
FROM dependencies-stage AS build-stage
COPY . .
RUN go install


#
## Release Stage
FROM alpine:${ALPINE_VERSION} as release-stage
ENV TZ Asia/Taipei

RUN apk update && apk --no-cache upgrade musl && apk add --no-cache tzdata curl bash lftp

COPY --from=build-stage /go/bin/go-clean-arch /usr/bin/

EXPOSE 5003

ENTRYPOINT ["go-clean-arch"]  
