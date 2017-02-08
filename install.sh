#!/bin/bash

set -eu

cd $(cd $(dirname ${BASH_SOURCE:0}); pwd)

cd cmd

go build cowsay.go
ln -fs cowsay cowthink

[ ! -d ~/bin ] && mkdir ~/bin

cp cowsay ~/bin/
cp cowthink ~/bin/

