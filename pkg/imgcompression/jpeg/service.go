package jpeg

import (
	"context"
	"github.com/bitstored/compression-service/pkg/errors"
	"image"
	"math"
)

type JPEGCompressor struct {
}

func NewCompressor() *JPEGCompressor {
	return &JPEGCompressor{}
}

func (c *JPEGCompressor) Compress(ctx context.Context, in image.Image, level interface{}) ([]byte, *errors.Err) {
	img := image2Array(in)
	ints, err := compressImage(img)
	bytes := make([]byte, 0)
	for i := range ints {
		bytes = append(bytes, byte(ints[i]))
	}
	return bytes, err
}

func (c *JPEGCompressor) Decompress(ctx context.Context, data []byte) (image.Image, *errors.Err) {
	ints := make([]int8, 0)
	for i := range data {
		ints = append(ints, int8(ints[i]))
	}
	arr, err := decompressImage(ints)
	if err != nil {
		return nil, err
	}
	img := arraytoImage(arr)
	return img, nil
}

func compressImage(in [][]RGBPixel) ([]int8, *errors.Err) {
	out, err := toYCbCr(in)
	if err != nil {
		return nil, err
	}

	out, err = resize(out, int(math.Ceil(float64(len(in))/8*8)), int(math.Ceil(float64(len(in[0]))/8*8)))
	if err != nil {
		return nil, err
	}

	out = toSigned(out)
	bytes := applyCosineAndEncode(out)
	// real size
	is := int2Int8Array(len(in))
	bytes = append(bytes, is...)
	is = int2Int8Array(len(in[0]))
	bytes = append(bytes, is...)

	return bytes, nil
}
func decompressImage(data []int8) ([][]RGBPixel, *errors.Err) {
	image, err := decodeEntropyAndApplyCosine(data)
	if err != nil {
		return nil, err
	}

	x := len(data) - 8
	y := len(data) - 4
	image, err = resize(image, int8Array2Int(data[x:x+4]), int8Array2Int(data[y:y+4]))
	if err != nil {
		return nil, err
	}

	image = toUnsigned(image)

	out, err := toRGB(image)
	if err != nil {
		return nil, err
	}
	return out, nil
}
