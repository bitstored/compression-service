package textcompression

import (
	"bytes"
	"compress/zlib"
	"context"
	"io/ioutil"

	"github.com/bitstored/compression-service/pkg/errors"
)

var dict = []byte("Hello, World!\n")

func (c *ZlibCompressor) Decompress(ctx context.Context, text []byte) ([]byte, *errors.Err) {
	b := bytes.NewReader(text)
	r, err := zlib.NewReader(b)
	defer r.Close()

	if err != nil {
		return nil, errors.NewError(errors.KindMalformedStructure, err.Error())
	}

	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errors.NewError(errors.KindMalformedStructure, err.Error())
	}
	return bytes, nil
}
