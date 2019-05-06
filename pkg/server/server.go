package server

import (
	"context"
	"github.com/bitstored/compression-service/pkg/imgcompression"
)

func CompressImage(ctx *context.Context, img [][]imgcompression.RGBPixel) []int8 {

	return nil
}

func DecompressImage(ctx *context.Context, img []int8) [][]imgcompression.RGBPixel {
	//img1 := toUnsigned(img)

	return nil
}

func CompressText(ctx *context.Context, text []byte) []byte {
	return nil
}

func DecompressText(ctx *context.Context, text []byte) []byte {
	return nil
}
