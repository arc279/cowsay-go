#!/bin//bash

set -eu

cd $(cd $(dirname ${BASH_SOURCE:0}); pwd)

for input in share/cows/*.cow; do
    output=share/cows-go/$(basename $input)
    echo "$input -> $output"
    sed \
        -e '/the_cow/d' \
        -e '/EOC/d' \
        -e '/^\(#.*\)$/d' \
        -e 's/\\\\/\\/g' \
        -e 's/$thoughts/{{.Thoughts}}/g' \
        -e 's/$eyes/{{.Eyes}}/g' \
        -e 's/$tongue/{{.Tongue}}/g' <$input >$output
done

cd share/cows-go
go-bindata -ignore=\\.gitignore -pkg cowsay -o ../../lib/bindata.go .
