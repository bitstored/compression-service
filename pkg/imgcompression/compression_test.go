package imgcompression

import (
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

var (
	tile1 = [][]YCbCrPixel{
		{YCbCrPixel{1, 2, 3}, YCbCrPixel{1, 2, 3}, YCbCrPixel{1, 2, 3}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{1, 2, 3}, YCbCrPixel{1, 2, 3}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{1, 2, 3}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
	}
	tile2 = [][]YCbCrPixel{
		{YCbCrPixel{1, 3, 5}, YCbCrPixel{9, 4, 45}, YCbCrPixel{-1, -2, -3}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{-45, -4, 3}, YCbCrPixel{0, -52, -43}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{1, 2, 93}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
	}
	tile3 = [][]YCbCrPixel{
		{YCbCrPixel{}, YCbCrPixel{9, 4, 45}, YCbCrPixel{-1, -2, -3}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{}, YCbCrPixel{0, -52, -43}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
	}
	tile4 = [][]YCbCrPixel{
		{YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
		{YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}, YCbCrPixel{}},
	}
	tile5 = [][]RGBPixel{
		{RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{22, 45, 67}, RGBPixel{22, 45, 67}, RGBPixel{22, 45, 67}, RGBPixel{22, 45, 67}, RGBPixel{22, 45, 67}},
		{RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{22, 45, 67}, RGBPixel{22, 45, 67}, RGBPixel{22, 45, 67}, RGBPixel{22, 45, 67}, RGBPixel{22, 45, 67}},
		{RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{22, 45, 67}, RGBPixel{22, 45, 67}, RGBPixel{22, 45, 67}, RGBPixel{22, 45, 67}, RGBPixel{22, 45, 67}},
		{RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{22, 45, 67}, RGBPixel{22, 45, 67}, RGBPixel{22, 45, 67}, RGBPixel{22, 45, 67}, RGBPixel{22, 45, 67}},
		{RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{22, 45, 67}, RGBPixel{22, 45, 67}, RGBPixel{72, 34, 167}, RGBPixel{72, 34, 167}, RGBPixel{72, 34, 167}},
		{RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{72, 34, 167}, RGBPixel{72, 34, 167}, RGBPixel{72, 4, 67}, RGBPixel{72, 4, 67}, RGBPixel{72, 4, 67}},
		{RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{72, 4, 67}, RGBPixel{72, 4, 67}, RGBPixel{72, 4, 67}, RGBPixel{72, 4, 67}, RGBPixel{2, 4, 67}},
		{RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{2, 4, 67}, RGBPixel{2, 4, 67}, RGBPixel{2, 4, 67}, RGBPixel{2, 4, 67}, RGBPixel{2, 4, 67}},
	}
	tile6 = [][]RGBPixel{
		{RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}},
		{RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}},
		{RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}},
		{RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}},
		{RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}},
		{RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}},
		{RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}},
		{RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}, RGBPixel{1, 2, 3}},
	}
	tile7 = [][]YCbCrPixel{
		{YCbCrPixel{1, 2, 3}, YCbCrPixel{1, 2, 3}, YCbCrPixel{1, 2, 3}, YCbCrPixel{22, 45, 67}, YCbCrPixel{22, 45, 67}, YCbCrPixel{22, 45, 67}, YCbCrPixel{22, 45, 67}, YCbCrPixel{22, 45, 67}},
		{YCbCrPixel{1, 2, 3}, YCbCrPixel{1, 2, 3}, YCbCrPixel{1, 2, 3}, YCbCrPixel{22, 45, 67}, YCbCrPixel{22, 45, 67}, YCbCrPixel{22, 45, 67}, YCbCrPixel{22, 45, 67}, YCbCrPixel{22, 45, 67}},
		{YCbCrPixel{1, 2, 3}, YCbCrPixel{1, 2, 3}, YCbCrPixel{1, 2, 3}, YCbCrPixel{22, 45, 67}, YCbCrPixel{22, 45, 67}, YCbCrPixel{22, 45, 67}, YCbCrPixel{22, 45, 67}, YCbCrPixel{22, 45, 67}},
		{YCbCrPixel{1, 2, 3}, YCbCrPixel{1, 2, 3}, YCbCrPixel{1, 2, 3}, YCbCrPixel{22, 45, 67}, YCbCrPixel{22, 45, 67}, YCbCrPixel{22, 45, 67}, YCbCrPixel{22, 45, 67}, YCbCrPixel{22, 45, 67}},
		{YCbCrPixel{1, 2, 3}, YCbCrPixel{1, 2, 3}, YCbCrPixel{1, 2, 3}, YCbCrPixel{22, 45, 67}, YCbCrPixel{22, 45, 67}, YCbCrPixel{72, 34, 167}, YCbCrPixel{72, 34, 167}, YCbCrPixel{72, 34, 167}},
		{YCbCrPixel{1, 2, 3}, YCbCrPixel{1, 2, 3}, YCbCrPixel{1, 2, 3}, YCbCrPixel{72, 34, 167}, YCbCrPixel{72, 34, 167}, YCbCrPixel{72, 4, 67}, YCbCrPixel{72, 4, 67}, YCbCrPixel{72, 4, 67}},
		{YCbCrPixel{1, 2, 3}, YCbCrPixel{1, 2, 3}, YCbCrPixel{1, 2, 3}, YCbCrPixel{72, 4, 67}, YCbCrPixel{72, 4, 67}, YCbCrPixel{72, 4, 67}, YCbCrPixel{72, 4, 67}, YCbCrPixel{2, 4, 67}},
		{YCbCrPixel{1, 2, 3}, YCbCrPixel{1, 2, 3}, YCbCrPixel{1, 2, 3}, YCbCrPixel{2, 4, 67}, YCbCrPixel{2, 4, 67}, YCbCrPixel{2, 4, 67}, YCbCrPixel{2, 4, 67}, YCbCrPixel{2, 4, 67}},
	}
)

func TestMultiply(t *testing.T) {
	in := RGBPixel{223, 145, 156}
	aux := in.multiply(RGB2YcbCrMatrix)
	out := aux.multiply(YcbCr2RGBMatrix)

	require.Truef(t, math.Abs(float64(in.R-out.R)) <= 1, "abs is : %f, value is %d, expected %d", math.Abs(float64(in.R-out.R)), out.R, in.R)
	require.Truef(t, math.Abs(float64(in.G-out.G)) <= 1, "abs is : %f, value is %d, expected %d", math.Abs(float64(in.G-out.G)), out.G, in.G)
	require.Truef(t, math.Abs(float64(in.B-out.B)) <= 1, "abs is : %f, value is %d, expected %d", math.Abs(float64(in.B-out.B)), out.B, in.B)
}

func TestToYCbCr(t *testing.T) {

	ts := []struct {
		Name  string
		Image [][]RGBPixel
		Error bool
	}{
		{
			"EmptyImage", [][]RGBPixel{}, true,
		},
		{
			"Ok", tile5, false,
		},
	}

	for _, tc := range ts {
		t.Run(tc.Name, func(t *testing.T) {
			rez, err := toYCbCr(tc.Image)
			if tc.Error {
				require.Error(t, err.Error())
			} else {
				require.NoError(t, err.Error())
				require.NotNil(t, rez)
			}
		})
	}
}

func TestExtraxt8x8Tile(t *testing.T) {
	image1 := make([][]YCbCrPixel, 16)
	image2 := make([][]YCbCrPixel, 3)

	for i := range image1 {
		image1[i] = make([]YCbCrPixel, 16)
	}
	for i := range image2 {
		image2[i] = make([]YCbCrPixel, 7)
	}

	ts := []struct {
		Name  string
		Image [][]YCbCrPixel
		X     int
		Y     int
		Error bool
	}{
		{
			"EmptyImage", [][]YCbCrPixel{}, 0, 0, true,
		},
		{
			"Negative Index", [][]YCbCrPixel{}, -10, 0, true,
		},
		{
			"Too big Index", [][]YCbCrPixel{}, 100000, 0, true,
		},
		{
			"Small Image index 0 0 ", image2, 0, 0, true,
		},
		{
			"Small Image index 2 2 ", image2, 2, 2, true,
		},
		{
			"OK Image index overflow ", image1, 10, 10, true,
		},
		{
			"Ok index start", image1, 0, 0, false,
		},
		{
			"Ok index middle", image1, 2, 4, false,
		},
		{
			"Ok index last", image1, 8, 8, false,
		},
	}
	for _, tc := range ts {
		t.Run(tc.Name, func(t *testing.T) {
			rez, err := extract8x8Tile(tc.Image, tc.X, tc.Y)
			if tc.Error {
				require.Error(t, err.Error())
			} else {
				require.NoError(t, err.Error())
				require.NotNil(t, rez)
				require.Equal(t, blockLen, len(rez))
				require.Equal(t, blockLen, len(rez[0]))
			}
		})
	}
}

func TestEncodeEntropy(t *testing.T) {
	ts := []struct {
		Name       string
		Tile       [][]YCbCrPixel
		ExpectedY  []int8
		ExpectedCb []int8
		ExpectedCr []int8
		HasError   bool
	}{
		{
			Name:       "EncodeEntropySameValue",
			Tile:       tile1,
			ExpectedY:  []int8{1, 1, 6},
			ExpectedCb: []int8{1, 2, 6},
			ExpectedCr: []int8{1, 3, 6},
			HasError:   false,
		},
		{
			Name:       "EncodeEntropyDifferentValues",
			Tile:       tile2,
			ExpectedY:  []int8{6, 1, 1, 9, 1, -45, 1, 1, 1, 0, 1, -1, 1},
			ExpectedCb: []int8{6, 3, 1, 4, 1, -4, 1, 2, 1, -52, 1, -2, 1},
			ExpectedCr: []int8{6, 5, 1, 45, 1, 3, 1, 93, 1, -43, 1, -3, 1},
			HasError:   false,
		},
		{
			Name:       "EncodeEntropyStartsWithZero",
			Tile:       tile3,
			ExpectedY:  []int8{4, 0, 1, 9, 1, 0, 3, -1, 1},
			ExpectedCb: []int8{5, 0, 1, 4, 1, 0, 2, -52, 1, -2, 1},
			ExpectedCr: []int8{5, 0, 1, 45, 1, 0, 2, -43, 1, -3, 1},
			HasError:   false,
		},
		{
			Name:       "EncodeEntropyFilledWithZero",
			Tile:       tile4,
			ExpectedY:  []int8{0},
			ExpectedCb: []int8{0},
			ExpectedCr: []int8{0},
			HasError:   false,
		},
		{
			Name:     "EncodeEntropyEmpty",
			Tile:     [][]YCbCrPixel{},
			HasError: true,
		},
	}

	for _, tc := range ts {
		t.Run(tc.Name, func(t *testing.T) {
			y, cb, cr, err := encodeEntropy(tc.Tile)
			if tc.HasError {
				require.Error(t, err.Error())
			} else {
				require.NoError(t, err.Error())
				require.EqualValues(t, tc.ExpectedY, y)
				require.EqualValues(t, tc.ExpectedCb, cb)
				require.EqualValues(t, tc.ExpectedCr, cr)
			}
		})
	}
}

func TestCompressImage(t *testing.T) {

	ts := []struct {
		Name     string
		Tile     [][]RGBPixel
		HasError bool
	}{
		{
			Name:     "Compress Ok diff",
			Tile:     tile5,
			HasError: false,
		},
		{
			Name:     "Compress Ok equal",
			Tile:     tile6,
			HasError: false,
		},
		{
			Name:     "Compress Empty",
			Tile:     [][]RGBPixel{},
			HasError: true,
		},
	}

	for _, tc := range ts {
		t.Run(tc.Name, func(t *testing.T) {
			rez, err := compressImage(tc.Tile)
			if tc.HasError {
				require.Error(t, err.Error())
			} else {
				require.NoError(t, err.Error())
				require.True(t, 8*8*3 > len(rez))
			}
		})
	}
}
