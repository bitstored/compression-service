package png

import (
	"bytes"
	"context"
	"image"
	"image/png"

	"github.com/bitstored/compression-service/pkg/errors"
)

func (c *PNGCompressor) Decompress(ctx context.Context, img []byte) (image.Image, *errors.Err) {
	writer := new(bytes.Buffer)
	_, err := writer.Write(img)
	if err != nil {
		return nil, errors.NewError(errors.KindMalformedStructure, err.Error())
	}
	imgout, err := png.Decode(writer)
	if err != nil {
		return nil, errors.NewError(errors.KindMalformedStructure, err.Error())
	}
	return imgout, nil
}
