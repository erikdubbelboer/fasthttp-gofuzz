#!/bin/sh

go-fuzz-build -o fuzz.zip github.com/erikdubbelboer/fasthttp-gofuzz/request
nice -n 15 go-fuzz -bin=fuzz.zip -workdir=. -procs=2 -testoutput=true > output.txt
