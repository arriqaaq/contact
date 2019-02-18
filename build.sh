#!/bin/bash
set -e

VERSION=$(cat VERSION)
echo "VERSION: $VERSION"

APP_IMG_NAME="flash/contact"
APP_IMG=${APP_IMG_NAME}:latest


DOCKERFILE_APP="Dockerfile"


echo Building app image: $APP_IMG
docker build -t $APP_IMG . -f $DOCKERFILE_APP

docker-compose up
