#!/bin/sh

i=0

for entry in './test/plugin-buildmode'/*; do
  [[ "$entry" != *.go ]] && continue

  go build -buildmode=plugin -o "./test/temp/${i}.so" "$entry"
  echo "$entry"

  i=(i+ 1)
done

go test -v ./test/...
