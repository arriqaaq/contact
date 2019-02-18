#!/bin/bash


echo Building production image: $IMG
# # tag it
git tag -af "$BUILD_VERSION" -m "version $VERSION"
git push --set-upstream origin master
git push -f --tags
# run build
./build.sh