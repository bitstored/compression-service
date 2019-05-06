package imgcompression

import (
	"context"
	"github.com/bitstored/compression-service/pkg/errors"
	"image"
)

type Compressor interface {
	Compress(ctx context.Context, img image.Image, iLevel interface{}) ([]byte, *errors.Err)
	Decompress(ctx context.Context, img []byte) (image.Image, *errors.Err)
}
