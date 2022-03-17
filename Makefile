# Copyright © 2022 Durudex

# This source code is licensed under the MIT license found in the
# LICENSE file in the root directory of this source tree.

.PHONY: build
build:
	go mod download && CGO_ENABLE=0 GOOS=linux go build -o ./.bin/app ./cmd/main.go

.PHONY: run
run: build
	docker-compose up --remove-orphans app

.PHONY: gqlgen
gqlgen:
	go get -d github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen generate --config ./gqlgen.yml
	go mod tidy

.DEFAULT_GOAL := run
