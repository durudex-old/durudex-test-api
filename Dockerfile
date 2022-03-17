# Copyright © 2022 Durudex

# This source code is licensed under the MIT license found in the
# LICENSE file in the root directory of this source tree.

FROM alpine:latest

WORKDIR /root/

ENTRYPOINT ["./app"]
