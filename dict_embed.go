//go:build go1.16 && !ne
// +build go1.16,!ne

package gse

import (
	"bytes"
	"compress/zlib"
	_ "embed"
	"io"
)

var (
	//go:embed data/dict/zh/t_1.txt
	zhT string

	//go:embed data/dict/zh/s_1.txt
	zhS string
)

//go:embed data/dict/zh/stop_tokens.txt
var stopDict string

func Decompressed(s string) string {
	var b bytes.Buffer
	b.WriteString(s)
	r, err := zlib.NewReader(&b)
	if err != nil {
		panic(err)
	}
	defer r.Close()
	decompressedData, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return string(decompressedData)
}
