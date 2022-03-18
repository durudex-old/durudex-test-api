# Copyright © 2022 Durudex

# This source code is licensed under the MIT license found in the
# LICENSE file in the root directory of this source tree.

FROM golang:1.18 AS builder

RUN go version

COPY . /github.com/durudex/durudex-test-api/
WORKDIR /github.com/durudex/durudex-test-api/

RUN go mod download
RUN CGO_ENABLE=0 GOOS=linux go build -o ./.bin/app ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/durudex/durudex-test-api/.bin/app .

CMD ["./app"]
