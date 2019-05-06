package jpeg

import (
	"encoding/binary"
	"github.com/bitstored/compression-service/pkg/errors"
	"image"
	"image/color"
	"math"
)

const (
	ratio               = image.YCbCrSubsampleRatio422
	blockLen            = 8
	transformationValue = 128
)

var (
	RGB2YcbCrMatrix = [][]float64{
		{0.299, 0.587, 0.114},
		{-0.168736, -0.331264, 0.5},
		{0.5, -0.418688, -0.081312},
	}
	YcbCr2RGBMatrix = [][]float64{
		{1, 0, 1.402},
		{1, -0.344136, -0.714136},
		{1, 1.772, 0},
	}
	quantizationMatrix = [][]int16{
		{16, 11, 10, 16, 24, 40, 51, 61},
		{12, 12, 14, 19, 26, 58, 60, 55},
		{14, 13, 16, 24, 40, 57, 69, 56},
		{14, 17, 22, 29, 51, 87, 80, 62},
		{18, 22, 37, 56, 68, 109, 103, 77},
		{24, 35, 55, 64, 81, 104, 113, 92},
		{49, 64, 78, 87, 103, 121, 120, 101},
		{72, 92, 95, 98, 112, 100, 103, 99},
	}
	chromianceMatrix = [][]int16{
		{17, 18, 24, 47, 99, 99, 99, 99},
		{18, 21, 26, 66, 99, 99, 99, 99},
		{24, 26, 56, 99, 99, 99, 99, 99},
		{47, 66, 99, 99, 99, 99, 99, 99},
		{99, 99, 99, 99, 99, 99, 99, 99},
		{99, 99, 99, 99, 99, 99, 99, 99},
		{99, 99, 99, 99, 99, 99, 99, 99},
		{99, 99, 99, 99, 99, 99, 99, 99},
	}
)

type Pair struct {
	Value int8
	Freq  int8
}

func maxInt16(a, b int16) int16 {
	if a > b {
		return a
	}
	return b
}

func minInt16(a, b int16) int16 {
	if a <= b {
		return a
	}
	return b
}

/* mul - multiplies a float64 by an int16
 * @returns an int16 in range of an uint8
 */
func mul(a float64, b int16) int16 {
	c := a * float64(b)
	d := math.Round(float64(c))
	return int16(d)
}

/* resize resizes a matrix to new dimensions x and y
 * @returns the resizesd matrix
 */
func resize(in [][]YCbCrPixel, x int, y int) ([][]YCbCrPixel, *errors.Err) {

	if in == nil || len(in) == 0 || len(in[0]) == 0 {
		return nil, errors.NewError(errors.KindEmptyImage, "image can not be resized, image is empty")
	}

	if x < 0 || y < 0 {
		return in, errors.NewError(errors.KindNotPermitedOperation, "image can not be resized to negative size")
	}

	if y > len(in) {
		n := y - len(in)
		for i := 0; i < n; i++ {
			row := make([]YCbCrPixel, len(in[0]))
			in = append(in, row)
		}
	} else if y < len(in) {
		in = in[0:y]
	}

	if x > len(in[0]) {
		for j := range in {
			aux := make([]YCbCrPixel, x-len(in))
			in[j] = append(in[j], aux...)
		}
	} else if x < len(in[0]) {
		for j := range in {
			in[j] = in[j][0:x]
		}
	}

	return in, nil
}

func int8Array2Int(a []int8) int {
	is := make([]byte, 0)
	for i := range a {
		is = append(is, byte(a[i]))
	}
	i := binary.LittleEndian.Uint32(is)
	return int(i)
}

func int2Int8Array(a int) []int8 {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(a))
	is := make([]int8, 0)

	for i := range bs {
		is = append(is, int8(bs[i]))
	}
	return is
}

func image2Array(img image.Image) [][]RGBPixel {
	b := img.Bounds()
	out := make([][]RGBPixel, b.Dx())
	for x := 0; x < b.Dx(); x++ {
		out[x] = make([]RGBPixel, b.Dy())
		for y := 0; y < b.Dy(); y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			out[x][y].R = int16(r)
			out[x][y].G = int16(g)
			out[x][y].B = int16(b)
		}
	}
	return out
}

func arraytoImage(img [][]RGBPixel) image.Image {
	out := image.NewNRGBA(image.Rect(0, 0, len(img), len(img[0])))

	for x := 0; x < len(img); x++ {
		for y := 0; y < len(img[0]); y++ {
			col := color.RGBA{uint8(img[x][y].R), uint8(img[x][y].G), uint8(img[x][y].B), 255}
			out.Set(x, y, col)
		}
	}

	return out
}
