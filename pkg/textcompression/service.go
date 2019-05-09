package textcompression

import (
	"context"
	"github.com/bitstored/compression-service/pb"
	"github.com/bitstored/compression-service/pkg/errors"
)

type Service interface {
	Compress(ctx context.Context, text []byte, level pb.CompressionLevel) ([]byte, *errors.Err)
	Decompress(ctx context.Context, text []byte) ([]byte, *errors.Err)
}

type ZlibCompressor struct {
}

func NewZlibCompressor() *ZlibCompressor {
	return &ZlibCompressor{}
}
