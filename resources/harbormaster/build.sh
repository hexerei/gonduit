#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

rm -r /home/hexerei/go/src

mkdir -p /home/hexerei/go/src/github.com/hexerei
ln -sf $(pwd) /home/hexerei/go/src/github.com/hexerei/
cd /home/hexerei/go/src/github.com/hexerei/gonduit

glide install
go build $(glide novendor)
go test $(glide novendor)
