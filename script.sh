#!/bin/bash

shopt -s globstar; for i in **/go.{mod,sum}; do
    echo "processing $i"
    awk "!/github\\.com\\/dev-monteiro-koho\\/monteiro-test\\/.*/ || /module/" "$i" > gomod.tmp
    mv gomod.tmp "$i"
done
