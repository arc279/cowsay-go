#!/bin//bash

set -eu

cd $(cd $(dirname ${BASH_SOURCE:0}); pwd)

SRC_DIR=${1:-share/cows/}
DST_DIR=~/.cowsay-go/

[ ! -d $DST_DIR ] && mkdir $DST_DIR

cat <<EOD >$DST_DIR/.uninstall.sh
#!/bin/bash
cd ~ && rm -rf $DST_DIR
EOD

function conv() {
    sed \
        -e 's/\\e/'$'\e/g' \
        -e 's/\\u/\\\\u/g' \
        -e 's/\\N{U+\([0-9]\+\)}/\\u\1/g' \
        -e 's/\\\\/\\/g' \
        -e '/binmode/d' \
        -e '/the_cow/d' \
        -e '/EOC/d' \
        -e '/^#.*$/d' \
        -e 's/$thoughts/{{.Thoughts}}/g' \
        -e 's/${thoughts}/{{.Thoughts}}/g' \
        -e 's/$eyes/{{.Eyes}}/g' \
        -e 's/${eyes}/{{.Eyes}}/g' \
        -e 's/$tongue/{{.Tongue}}/g' \
        -e 's/${tongue}/{{.Tongue}}/g'
}

PATTERN="${SRC_DIR}/*.cow"
for input in ${PATTERN}; do
    output=${DST_DIR}$(basename "$input")
    echo "$input -> $output"
    cat "$input" | conv | xargs -0 printf >"$output"
done

