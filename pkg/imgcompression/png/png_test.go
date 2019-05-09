package png

import (
	"context"
	"github.com/stretchr/testify/require"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"testing"
	"time"

	"github.com/bitstored/compression-service/pkg/imgcompression"
)

const (
	smallSizeX  = 10
	smallSizeY  = 12
	mediumSizeX = 40
	mediumSizeY = 50
	bigSizeX    = 400
	bigSizeY    = 500
	hugeSizeX   = 800
	hugeSizeY   = 1000
)

func TestCompressImage(t *testing.T) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var s imgcompression.Compressor = NewCompressor()

	imgSmallZiseSmallEntropy := image.NewNRGBA(image.Rect(0, 0, smallSizeX, smallSizeY))
	bounds := imgSmallZiseSmallEntropy.Bounds()
	for i := 0; i < bounds.Dx(); i++ {
		for j := 0; j < bounds.Dy(); j++ {
			imgSmallZiseSmallEntropy.Set(i, j, color.RGBA{uint8(i), uint8(j), uint8(i), 255})
		}
	}
	imgSmallSizeBigEntropy := image.NewNRGBA(image.Rect(0, 0, smallSizeX, smallSizeY))
	bounds = imgSmallSizeBigEntropy.Bounds()
	for i := 0; i < bounds.Dx(); i++ {
		for j := 0; j < bounds.Dy(); j++ {
			r := r1.Intn(255)
			g := r1.Intn(255)
			b := r1.Intn(255)
			imgSmallSizeBigEntropy.Set(i, j, color.RGBA{uint8(r), uint8(g), uint8(b), 255})
		}
	}

	imgMediumZiseSmallEntropy := image.NewNRGBA(image.Rect(0, 0, mediumSizeX, mediumSizeY))
	bounds = imgMediumZiseSmallEntropy.Bounds()
	for i := 0; i < bounds.Dx(); i++ {
		for j := 0; j < bounds.Dy(); j++ {
			imgMediumZiseSmallEntropy.Set(i, j, color.RGBA{uint8(i), uint8(j), uint8(i), 255})
		}
	}
	imgMediumSizeBigEntropy := image.NewNRGBA(image.Rect(0, 0, mediumSizeX, mediumSizeY))
	bounds = imgMediumSizeBigEntropy.Bounds()
	for i := 0; i < bounds.Dx(); i++ {
		for j := 0; j < bounds.Dy(); j++ {
			r := r1.Intn(255)
			g := r1.Intn(255)
			b := r1.Intn(255)
			imgMediumSizeBigEntropy.Set(i, j, color.RGBA{uint8(r), uint8(g), uint8(b), 255})
		}
	}

	imgBigZiseSmallEntropy := image.NewNRGBA(image.Rect(0, 0, bigSizeX, bigSizeY))
	bounds = imgBigZiseSmallEntropy.Bounds()
	for i := 0; i < bounds.Dx(); i++ {
		for j := 0; j < bounds.Dy(); j++ {
			imgBigZiseSmallEntropy.Set(i, j, color.RGBA{uint8(i), uint8(j), uint8(i), 255})
		}
	}
	imgBigSizeBigEntropy := image.NewNRGBA(image.Rect(0, 0, bigSizeX, bigSizeY))
	bounds = imgBigSizeBigEntropy.Bounds()
	for i := 0; i < bounds.Dx(); i++ {
		for j := 0; j < bounds.Dy(); j++ {
			r := r1.Intn(255)
			g := r1.Intn(255)
			b := r1.Intn(255)
			imgBigSizeBigEntropy.Set(i, j, color.RGBA{uint8(r), uint8(g), uint8(b), 255})
		}
	}
	imgHugeZiseSmallEntropy := image.NewNRGBA(image.Rect(0, 0, hugeSizeX, hugeSizeY))
	bounds = imgHugeZiseSmallEntropy.Bounds()
	for i := 0; i < bounds.Dx(); i++ {
		for j := 0; j < bounds.Dy(); j++ {
			imgHugeZiseSmallEntropy.Set(i, j, color.RGBA{uint8(i), uint8(j), uint8(i), 255})
		}
	}
	imgHugeSizeBigEntropy := image.NewNRGBA(image.Rect(0, 0, hugeSizeX, hugeSizeY))
	bounds = imgHugeSizeBigEntropy.Bounds()
	for i := 0; i < bounds.Dx(); i++ {
		for j := 0; j < bounds.Dy(); j++ {
			r := r1.Intn(255)
			g := r1.Intn(255)
			b := r1.Intn(255)
			imgHugeSizeBigEntropy.Set(i, j, color.RGBA{uint8(r), uint8(g), uint8(b), 255})
		}
	}

	ts := []struct {
		Name        string
		Image       image.Image
		SizeBest    int
		SizeDefault int
		SizeNo      int
		SizeSpeed   int
		TimeBest    int64
		TimeDefault int64
		TimeNo      int64
		TimeSpeed   int64
	}{
		{
			Name:  "TestSmallImageSmallEnthropy",
			Image: imgSmallZiseSmallEntropy,
		},
		{
			Name:  "TestSmallImageBigEnthropy",
			Image: imgSmallSizeBigEntropy,
		},
		{
			Name:  "TestMediumImageSmallEnthropy",
			Image: imgMediumZiseSmallEntropy,
		},
		{
			Name:  "TestMediumImageBigEnthropy",
			Image: imgMediumSizeBigEntropy,
		},
		{
			Name:  "TestBigImageSmallEnthropy",
			Image: imgBigZiseSmallEntropy,
		},
		{
			Name:  "TestBigImageBigEnthropy",
			Image: imgBigSizeBigEntropy,
		},
		{
			Name:  "TestHugeImageSmallEnthropy",
			Image: imgHugeZiseSmallEntropy,
		},
		{
			Name:  "TestHugeImageHugeEnthropy",
			Image: imgHugeSizeBigEntropy,
		},
	}
	for _, tc := range ts {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()

			// Test with best compression
			start := time.Now()
			bytes, err := s.Compress(ctx, tc.Image, png.BestCompression)
			end := time.Now()
			require.NoError(t, err.Error())

			imag, err := s.Decompress(context.Background(), bytes)
			require.NoError(t, err.Error())

			testEqual(t, tc.Image, imag)
			tc.SizeBest = len(bytes)
			tc.TimeBest = end.Sub(start).Nanoseconds()

			// Test with default compression
			start = time.Now()
			bytes, err = s.Compress(ctx, tc.Image, png.DefaultCompression)
			end = time.Now()

			require.NoError(t, err.Error())

			imag, err = s.Decompress(context.Background(), bytes)
			require.NoError(t, err.Error())

			testEqual(t, tc.Image, imag)
			tc.SizeDefault = len(bytes)
			tc.TimeDefault = end.Sub(start).Nanoseconds()

			// Test with speed compression
			start = time.Now()
			bytes, err = s.Compress(ctx, tc.Image, png.BestSpeed)
			end = time.Now()

			require.NoError(t, err.Error())

			imag, err = s.Decompress(context.Background(), bytes)
			require.NoError(t, err.Error())

			testEqual(t, tc.Image, imag)
			tc.SizeSpeed = len(bytes)
			tc.TimeSpeed = end.Sub(start).Nanoseconds()

			// Test with no compression
			start = time.Now()
			bytes, err = s.Compress(ctx, tc.Image, png.NoCompression)
			end = time.Now()

			require.NoError(t, err.Error())

			imag, err = s.Decompress(context.Background(), bytes)
			require.NoError(t, err.Error())

			testEqual(t, tc.Image, imag)
			tc.SizeNo = len(bytes)
			tc.TimeNo = end.Sub(start).Nanoseconds()
		})
	}
}

func testEqual(t *testing.T, initial image.Image, resulted image.Image) {
	t.Helper()
	require.Equal(t, initial.Bounds().Dx(), resulted.Bounds().Dx())
	require.Equal(t, initial.Bounds().Dy(), resulted.Bounds().Dy())
	for x := 0; x < smallSizeX; x++ {
		for y := 0; y < smallSizeY; y++ {
			r1, g1, b1, _ := initial.At(x, y).RGBA()
			r2, g2, b2, _ := resulted.At(x, y).RGBA()
			require.EqualValues(t, r1, r2)
			require.EqualValues(t, g1, g2)
			require.EqualValues(t, b1, b2)
		}
	}
}
