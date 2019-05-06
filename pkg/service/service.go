package service

import (
	"context"
	"image"
	"image/png"

	"github.com/bitstored/compression-service/pb"
	"github.com/bitstored/compression-service/pkg/imgcompression"
	_ "github.com/bitstored/compression-service/pkg/imgcompression/jpeg"
	comp_png "github.com/bitstored/compression-service/pkg/imgcompression/png"
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
		//c = comp_jpeg.NewCompressor()
	}
	out, err := c.Compress(ctx, img, level)

	return out, err.Error()
}

func (s *Service) DecompressImage(ctx context.Context, img []byte, level png.CompressionLevel, imgType pb.ImageType) (image.Image, error) {
	var c imgcompression.Compressor

	if imgType == pb.ImageType_PNG {
		c = comp_png.NewCompressor()
	} else {
		//c = comp_jpeg.NewCompressor()
	}
	out, err := c.Decompress(ctx, img)

	return out, err.Error()
}

func (s *Service) CompressText(ctx context.Context, text []byte) ([]byte, error) {
	return nil, nil
}
func (s *Service) DecompressText(ctx context.Context, text []byte) ([]byte, error) {
	return nil, nil
}
