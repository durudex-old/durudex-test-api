# Copyright © 2022 Durudex
#
# This source code is licensed under the MIT license found in the
# LICENSE file in the root directory of this source tree.

.PHONY: build
build:
	docker build -f nginx/Dockerfile -t durudex-test-proxy:latest .

.PHONY: run
run: build
	docker-compose up -d --remove-orphans api
	docker-compose up -d --remove-orphans nginx

.DEFAULT_GOAL := run
