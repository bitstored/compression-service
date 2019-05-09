package service

import (
	"context"
	"image"
	"image/png"

	"github.com/bitstored/compression-service/pb"
	"github.com/bitstored/compression-service/pkg/imgcompression"
	comp_jpeg "github.com/bitstored/compression-service/pkg/imgcompression/jpeg"
	comp_png "github.com/bitstored/compression-service/pkg/imgcompression/png"
	"github.com/bitstored/compression-service/pkg/textcompression"
)

type Service struct {
}

func NewCompressionService() *Service {
	return &Service{}
}

func (s *Service) CompressImage(ctx context.Context, img image.Image, level png.CompressionLevel, imgType pb.ImageType) ([]byte, error) {
	var c imgcompression.Compressor

	if imgType == pb.ImageType_PNG {
		c = comp_png.NewCompressor()
	} else {
		c = comp_jpeg.NewCompressor()
	}
	out, err := c.Compress(ctx, img, level)

	if err != nil {
		return nil, err.Error()
	}
	return out, nil
}

func (s *Service) DecompressImage(ctx context.Context, img []byte, level png.CompressionLevel, imgType pb.ImageType) (image.Image, error) {
	var c imgcompression.Compressor

	if imgType == pb.ImageType_PNG {
		c = comp_png.NewCompressor()
	} else {
		c = comp_jpeg.NewCompressor()
	}
	out, err := c.Decompress(ctx, img)
	if err != nil {
		return nil, err.Error()
	}
	return out, nil
}

func (s *Service) CompressText(ctx context.Context, text []byte, l pb.CompressionLevel) ([]byte, error) {
	c := textcompression.NewZlibCompressor()
	out, err := c.Compress(ctx, text, l)
	if err != nil {
		return nil, err.Error()
	}
	return out, nil
}

func (s *Service) DecompressText(ctx context.Context, text []byte) ([]byte, error) {
	c := textcompression.NewZlibCompressor()
	out, err := c.Decompress(ctx, text)
	if err != nil {
		return nil, err.Error()
	}
	return out, nil
}
