package fuzz

import (
	"bufio"
	"bytes"

	"github.com/valyala/fasthttp"
)

func Fuzz(data []byte) int {
	res := fasthttp.AcquireResponse()
	if err := res.Read(bufio.NewReader(bytes.NewReader(data))); err != nil {
		return 0
	}

	w := bytes.Buffer{}
	if _, err := res.WriteTo(&w); err != nil {
		return 0
	}

	fasthttp.ReleaseResponse(res)

	return 1
}
