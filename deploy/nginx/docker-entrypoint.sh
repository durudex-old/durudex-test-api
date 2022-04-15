#!/usr/bin/env sh

# Copyright © 2022 Durudex

# This source code is licensed under the MIT license found in the
# LICENSE file in the root directory of this source tree.

set -eu

envsubst '${SERVER_NAME}' < /nginx.conf.template > /etc/nginx/nginx.conf

exec "$@"
