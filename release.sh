#!/bin/bash

set -e
set -u

export VERSION="v2017102701"

# UI
pushd ui

rm -rf dist
mkdir -p dist
npm install
NODE_ENV=production npm run build
cp robots.txt dist/
mkdir -p dist/bundle
mv dist/index.html dist/robots.txt dist/static dist/bundle

popd

# dist
rm -rf dist

binfs ui/dist/bundle > binfs.out.go

mkdir -p dist/bastion-$VERSION
mkdir -p dist/blackbox-$VERSION
cp config.sample.toml dist/bastion-$VERSION/

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
