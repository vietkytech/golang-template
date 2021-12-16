GO_VERSION=1.15.7
DOCKER_VERSION=latest
DIST_NAME=golang-template

dev:
	docker run --rm -it --name dev \
		-p 8080:8080 \
		-v `pwd`/${DIST_NAME}:/go/src/github.com/vietkytech/${DIST_NAME}/${DIST_NAME} docker.chotot.org/${DIST_NAME}-dev:latest sh

build_dev:
	docker build -t docker.chotot.org/${DIST_NAME}-dev:latest -f Dockerfile.builder .


test:
	docker build -t docker.chotot.org/${DIST_NAME}:latest -f Dockerfile .
	docker run -it --rm --name test docker.chotot.org/${DIST_NAME}:latest sh

init-project:
	docker run --rm -v `pwd`/:/app/  debian:buster-20210816-slim cd /app/ && bash init.sh