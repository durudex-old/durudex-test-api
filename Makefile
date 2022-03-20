# Copyright © 2022 Durudex

# This source code is licensed under the MIT license found in the
# LICENSE file in the root directory of this source tree.

.PHONY: build
build:
	docker build -t durudex-test-api .

.PHONY: run
run: build
	docker-compose up --remove-orphans app

.PHONY: gqlgen
gqlgen:
	go get -d github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen generate --config ./gqlgen.yml
	go mod tidy

.DEFAULT_GOAL := run
