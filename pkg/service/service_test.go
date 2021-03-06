package service

import (
	"context"
	"github.com/bitstored/compression-service/pb"
	"github.com/stretchr/testify/require"
	"image"
	"image/color"
	"image/png"
	"testing"
)

func TestCompressImage(t *testing.T) {
	img := image.NewNRGBA(image.Rect(0, 0, 100, 100))
	bounds := img.Bounds()
	for i := 0; i < bounds.Dx(); i++ {
		for j := 0; j < bounds.Dy(); j++ {
			img.Set(i, j, color.RGBA{uint8(i), uint8(j), uint8(i), 255})
		}
	}
	s := NewCompressionService()
	ts := []struct {
		Name  string
		Level png.CompressionLevel
		Image image.Image
		Bytes []byte
	}{
		{
			Name:  "TestNoCpompression",
			Level: png.NoCompression,
			Image: img,
		},
		{
			Name:  "TestDefaultCompression",
			Level: png.DefaultCompression,
			Image: img,
		},
		{
			Name:  "TestBestSpeed",
			Level: png.BestSpeed,
			Image: img,
		},
		{
			Name:  "TestBestCompression",
			Level: png.BestCompression,
			Image: img,
		},
	}
	for _, tc := range ts {
		t.Run(tc.Name, func(t *testing.T) {
			bytes, err := s.CompressImage(context.Background(), tc.Image, tc.Level, pb.ImageType_PNG)
			require.NoError(t, err)
			tc.Bytes = bytes
			require.True(t, len(bytes) < 4*bounds.Dy()*bounds.Dx())
			imag, err := s.DecompressImage(context.Background(), tc.Bytes, tc.Level, pb.ImageType_PNG)
			require.NoError(t, err)
			for i := 0; i < bounds.Dx(); i++ {
				for j := 0; j < bounds.Dy(); j++ {
					t.Run("Test", func(t *testing.T) {
						col1 := imag.At(i, j)
						col2 := tc.Image.At(i, j)
						require.EqualValues(t, col1, col2)
					})
				}
			}
		})
	}
}
