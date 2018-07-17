# zenform - Provisioning tool for Zendesk
#
#   (C) 2018- nukosuke <nukosuke@lavabit.com>
#
# This software is released under MIT License.
# See LICENSE.

zenform: config/*.go command/*.go
	go build .
docker-image: Dockerfile zenform
	docker build -t xflagstudio/zenform .
install: zenform
	cp zenform /usr/local/bin/
clean:
	rm -f zenform
