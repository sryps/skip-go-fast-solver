#!/bin/sh

echo "Generating proto code"

proto_dirs=$(find . -path -prune -o -name 'buf.gen.yaml' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  cd $dir
  buf generate
  cd -
done
