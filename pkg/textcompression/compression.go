package textcompression

import (
	"bytes"
	"compress/zlib"
	"context"

	"github.com/bitstored/compression-service/pb"
	"github.com/bitstored/compression-service/pkg/errors"
)

func (c *ZlibCompressor) Compress(ctx context.Context, text []byte, l pb.CompressionLevel) ([]byte, *errors.Err) {
	var b bytes.Buffer
	level := compressionLevel(l)
	wld, err := zlib.NewWriterLevel(&b, level)
	if err != nil {
		return nil, errors.NewError(errors.KindMalformedStructure, err.Error())
	}
	wld.Write(text)
	wld.Close()
	return b.Bytes(), nil
}
