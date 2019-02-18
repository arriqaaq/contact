NAME   := flash/contact
LATEST := ${NAME}:latest
STABLE := ${NAME}:stable
VER := $(shell cat VERSION)
VERSION := ${NAME}:${VER}


.PHONY: build

build:
	./build.sh prod

test:
	./build.sh test

push:
	gcloud docker -- push ${STABLE}
	gcloud docker -- push ${LATEST}
	gcloud docker -- push ${VERSION}
