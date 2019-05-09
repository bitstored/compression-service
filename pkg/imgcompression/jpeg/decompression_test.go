package jpeg

import (
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestMultiply2(t *testing.T) {
	in := YCbCrPixel{223, 145, 156}
	aux := in.multiply(YcbCr2RGBMatrix)
	out := aux.multiply(RGB2YcbCrMatrix)

	require.Truef(t, math.Abs(float64(in.Y-out.Y)) <= 2, "Y abs is : %f, value is %d, expected %d", math.Abs(float64(in.Y-out.Y)), out.Y, in.Y)
	require.Truef(t, math.Abs(float64(in.Cb-out.Cb)) <= 2, "Cb abs is : %f, value is %d, expected %d", math.Abs(float64(in.Cb-out.Cb)), out.Cb, in.Cb)
	require.Truef(t, math.Abs(float64(in.Cr-out.Cr)) <= 3, "Cr abs is : %f, value is %d, expected %d", math.Abs(float64(in.Cr-out.Cr)), out.Cr, in.Cr)
}

func TestToRGB(t *testing.T) {

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
			out, err := toRGB(rez)
			if tc.Error {
				require.Error(t, err.Error())
			} else {
				require.NoError(t, err.Error())
				require.NotNil(t, out)
			}
			for i := range out {
				for j := range out[0] {
					require.True(t, math.Abs(float64(tc.Image[i][j].R-out[i][j].R)) <= 2)
					require.True(t, math.Abs(float64(tc.Image[i][j].G-out[i][j].G)) <= 2)
					require.True(t, math.Abs(float64(tc.Image[i][j].B-out[i][j].B)) <= 2)

				}
			}
		})
	}
}

func TestDecodeEntropyTile(t *testing.T) {

	ts := []struct {
		Name     string
		Expected [][]YCbCrPixel
		Slice    []int8
		HasError bool
	}{
		{
			Name:     "EncodeEntropySameValue",
			Expected: tile1,
			Slice:    []int8{1, 1, 6, 1, 2, 6, 1, 3, 6},
			HasError: false,
		},
		{
			Name:     "EncodeEntropyNil",
			Slice:    nil,
			Expected: nil,
			HasError: true,
		},
	}
	for _, tc := range ts {
		t.Run(tc.Name, func(t *testing.T) {
			tile, err := decodeEntropyTile(tc.Slice)
			if !tc.HasError {
				require.NoError(t, err.Error())
				for i := 0; i < blockLen; i++ {
					for j := 0; j < blockLen; j++ {
						require.EqualValuesf(t, tc.Expected[i][j].Y, tile[i][j].Y, "i : %d, j : %d\n", i, j)
						require.EqualValuesf(t, tc.Expected[i][j].Cb, tile[i][j].Cb, "i : %d, j : %d\n", i, j)
						require.EqualValuesf(t, tc.Expected[i][j].Cr, tile[i][j].Cr, "i : %d, j : %d\n", i, j)
					}
				}
			} else {
				require.Error(t, err.Error())
			}
		})
	}
}

func TestCosineTransform(t *testing.T) {
	ts := []struct {
		Name string
		Tile [][]YCbCrPixel
	}{
		{
			"TestIdenticalPyramidOK",
			tile1,
		},
		{
			"TestDifferentPyramidOK",
			tile2,
		},
		{
			"TestRandomOK",
			tile3,
		},
		{
			"TestZerosOK",
			tile4,
		},
		{
			"TestFull",
			tile7,
		},
	}
	for _, tc := range ts {
		t.Run(tc.Name, func(t *testing.T) {
			aux := funcDCT(tc.Tile)
			tile := funcICT(aux)
			for i := 0; i < blockLen; i++ {
				for j := 0; j < blockLen; j++ {
					t.Run("test", func(t *testing.T) {
						require.Truef(t, math.Abs(float64(tc.Tile[i][j].Y-tile[i][j].Y)) <= 2, "Y i : [%d] , j : [%d]  %d != %d\n", i, j, tc.Tile[i][j].Y, tile[i][j].Y)
						require.Truef(t, math.Abs(float64(tc.Tile[i][j].Cb-tile[i][j].Cb)) <= 2, "Cb i : [%d] , j : [%d]  %d != %d\n", i, j, tc.Tile[i][j].Cb, tile[i][j].Cb)
						require.Truef(t, math.Abs(float64(tc.Tile[i][j].Cr-tile[i][j].Cr)) <= 2, "Cr i : [%d], j : [%d]  %d != %d\n", i, j, tc.Tile[i][j].Cr, tile[i][j].Cr)
					})
				}
			}
		})
	}
}
func TestDecompressImage(t *testing.T) {

	ts := []struct {
		Name     string
		Tile     [][]RGBPixel
		HasError bool
	}{
		{
			Name:     "Compress Empty",
			Tile:     [][]RGBPixel{},
			HasError: true,
		},
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
	}

	for _, tc := range ts {
		t.Run(tc.Name, func(t *testing.T) {
			rez, err := compressImage(tc.Tile)
			if tc.HasError {
				require.Error(t, err.Error())
			} else {
				require.NoError(t, err.Error())
				_, err := decompressImage(rez)
				require.NoError(t, err.Error())
				// for i := 0; i < blockLen; i++ {
				// 	for j := 0; j < blockLen; j++ {
				// 		t.Run("test", func(t *testing.T) {
				// 			require.Truef(t, math.Abs(float64(tc.Tile[i][j].R-tile[i][j].R)) <= 2, "R i : [%d] , j : [%d]  %d != %d\n", i, j, tc.Tile[i][j].R, tile[i][j].R)
				// 			require.Truef(t, math.Abs(float64(tc.Tile[i][j].G-tile[i][j].G)) <= 2, "G i : [%d] , j : [%d]  %d != %d\n", i, j, tc.Tile[i][j].G, tile[i][j].G)
				// 			require.Truef(t, math.Abs(float64(tc.Tile[i][j].B-tile[i][j].B)) <= 2, "B i : [%d], j : [%d]  %d != %d\n", i, j, tc.Tile[i][j].B, tile[i][j].B)
				// 		})
				// 	}
				// }
			}
		})
	}
}
