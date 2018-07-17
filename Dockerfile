# zenform - Provisioning tool for Zendesk
#
#   (C) 2018- nukosuke <nukosuke@lavabit.com>
#
# This software is released under MIT License.
# See LICENSE.

FROM golang:1.9-alpine AS build
WORKDIR /go/src/github.com/xflagstudio/zenform
COPY . .
RUN \
    apk add --update git; \
    go get -u github.com/golang/dep/cmd/dep; \
    dep ensure; \
    go build .


FROM alpine
MAINTAINER nukosuke <nukosuke@lavabit.com>
COPY --from=build /go/src/github.com/xflagstudio/zenform/zenform /usr/bin/zenform
ENTRYPOINT ["/usr/bin/zenform"]
