GO_VERSION=1.15.7
DOCKER_VERSION=latest
DIST_NAME=multi-rejected-reasons

dev:
	docker run --rm -it --name dev \
		-p 8080:8080 \
		-v `pwd`/${DIST_NAME}:/go/src/git.chotot.org/fse/${DIST_NAME}/${DIST_NAME} docker.chotot.org/${DIST_NAME}-dev:latest sh

build_dev:
	docker build -t docker.chotot.org/${DIST_NAME}-dev:latest -f Dockerfile.builder .


test:
	docker build -t docker.chotot.org/${DIST_NAME}:latest -f Dockerfile .
	docker run -it --rm --name test docker.chotot.org/${DIST_NAME}:latest sh
