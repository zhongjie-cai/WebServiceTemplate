#!/bin/bash
export GO_VERSION=1.15
export GO_DOWNLOAD_URL=https://storage.googleapis.com/golang/go$GO_VERSION.linux-amd64.tar.gz

export GOPATH=$HOME/go
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

sudo mkdir ${GOPATH}
sudo chown ${USER} -R ${GOPATH}

sudo apt update --fix-missing && apt upgrade -y
sudo apt install --no-install-recommends -y gcc

wget "$GO_DOWNLOAD_URL" -O golang.tar.gz
tar -zxvf golang.tar.gz
sudo mv go ${GOROOT}

go version
