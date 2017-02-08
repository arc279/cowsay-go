#!/bin//bash

set -eu

cd $(cd $(dirname ${BASH_SOURCE:0}); pwd)

COWS_DIR=~/.cowsay-go/
[ ! -d $COWS_DIR ] && mkdir $COWS_DIR

cat <<EOD >$COWS_DIR/uninstall.sh
#!/bin/bash
cd ~
rm -rf $COWS_DIR
EOD

for input in share/cows/*.cow; do
    output=${COWS_DIR}$(basename $input)
    echo "$input -> $output"
    sed \
        -e '/the_cow/d' \
        -e '/EOC/d' \
        -e '/^#.*$/d' \
        -e 's/\\\\/\\/g' \
        -e 's/$thoughts/{{.Thoughts}}/g' \
        -e 's/${thoughts}/{{.Thoughts}}/g' \
        -e 's/$eyes/{{.Eyes}}/g' \
        -e 's/${eyes}/{{.Eyes}}/g' \
        -e 's/$tongue/{{.Tongue}}/g' \
        -e 's/${tongue}/{{.Tongue}}/g' \
        <$input >$output
done

