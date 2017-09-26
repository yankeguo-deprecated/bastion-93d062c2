#!/bin/bash

set -e
set -u

export VERSION="v2017092601"

# UI
pushd ui

rm -rf dist
mkdir -p dist
npm install
NODE_ENV=production npm run build
cp robots.txt dist/

popd

# dist
rm -rf dist/bastion-$VERSION*
rm -rf dist/blackbox-$VERSION*

mkdir -p dist/bastion-$VERSION
mkdir -p dist/blackbox-$VERSION
mkdir -p dist/bastion-$VERSION/public
cp -r ui/dist/* dist/bastion-$VERSION/public/
cp config.sample.toml dist/bastion-$VERSION/public/

# bastion
GOOS=linux GOARCH=amd64 go build -o dist/bastion-$VERSION/bastion

# blackbox
pushd blackbox
GOOS=linux GOARCH=amd64 go build -o ../dist/blackbox-$VERSION/blackbox
popd

# package
pushd dist
zip -r bastion-$VERSION-linux-amd64.zip bastion-$VERSION
gpg -abs bastion-$VERSION-linux-amd64.zip
zip -r blackbox-$VERSION-linux-amd64.zip blackbox-$VERSION
gpg -abs blackbox-$VERSION-linux-amd64.zip
popd
