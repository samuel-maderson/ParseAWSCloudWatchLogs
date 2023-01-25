package libs

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

func GzipDecompress(data string) string {
	reader, err := gzip.NewReader(bytes.NewReader([]byte(data)))
	if err != nil {
		panic(err)
	}

	defer reader.Close()

	decompressed, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	return string(decompressed)
}
