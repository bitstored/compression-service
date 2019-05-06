package png

import (
	"bytes"
	"context"
	"github.com/bitstored/compression-service/pkg/errors"
	"image"
	"image/png"
)

func (c *PNGCompressor) Compress(ctx context.Context, img image.Image, iLevel interface{}) ([]byte, *errors.Err) {
	level := iLevel.(png.CompressionLevel)

	writer := new(bytes.Buffer)
	encoder := png.Encoder{CompressionLevel: level}
	err := encoder.Encode(writer, img)

	if err != nil {
		return nil, errors.NewError(errors.KindMalformedStructure, err.Error())
	}

	return writer.Bytes(), nil
}
