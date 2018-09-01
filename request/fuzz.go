package fuzz

import (
	"bufio"
	"bytes"

	"github.com/valyala/fasthttp"
)

func Fuzz(data []byte) int {
	req := fasthttp.AcquireRequest()
	if err := req.Read(bufio.NewReader(bytes.NewReader(data))); err != nil {
		return 0
	}

	w := bytes.Buffer{}
	if _, err := req.WriteTo(&w); err != nil {
		return 0
	}

	return 1
}
