package jpeg

import (
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestMul(t *testing.T) {
	ts := []struct {
		Name    string
		Float32 float64
		Int16   int16
		Result  int16
	}{
		{
			"FloatZeroResultZero", 0.0, 123, 0,
		},
		{
			"IntZeroResultZero", 2.0, 0, 0,
		},
		{
			"OverflowTooBig", 2.0, 200, 400,
		},
		{
			"Overflow too small", -2.0, 23, -46,
		},
		{
			"OK", 2.0, 4, 8,
		},
	}
	for _, tc := range ts {
		rez := mul(tc.Float32, tc.Int16)
		require.EqualValues(t, tc.Result, rez)
	}
}

func TestInt8Array2Int(t *testing.T) {
	ts := []struct {
		Name  string
		Array []int8
		Value int
	}{
		{Name: "Test small", Array: []int8{0, 0, 0, 0}, Value: 0},
		{Name: "Test big -1", Array: []int8{-1, 0, 0, 0}, Value: 255},
		{Name: "Test big -128", Array: []int8{-128, 0, 0, 0}, Value: 128},
		{Name: "Test big -2", Array: []int8{-2, 0, 0, 0}, Value: 254},
		{Name: "Test big 127", Array: []int8{127, 0, 0, 0}, Value: 127},
		{Name: "Test number 1 poz", Array: []int8{12, 0, 0, 0}, Value: 12},
		{Name: "Test number 2 poz", Array: []int8{1, 1, 0, 0}, Value: 257},
		{Name: "Test number 3 poz", Array: []int8{1, 1, 1, 0}, Value: int(math.Pow(2, 16)) + int(math.Pow(2, 8)) + 1},
		{Name: "Test number 4 poz", Array: []int8{1, 0, 0, 1}, Value: int(math.Pow(2, 24)) + 1},
	}
	for _, tc := range ts {
		t.Run(tc.Name, func(t *testing.T) {
			rezValue := int8Array2Int(tc.Array)
			rezArray := int2Int8Array(tc.Value)
			require.EqualValuesf(t, tc.Array, rezArray, "not equal %v %v", tc.Array, rezArray)
			require.EqualValuesf(t, tc.Value, rezValue, "not equal %d %d", tc.Value, rezValue)

		})
	}
}

func TestRound(t *testing.T) {
	ts := []struct {
		Float float64
		Int   int16
	}{
		{
			0.5,
			1,
		},
		{
			0.55,
			1,
		},
		{
			0.45,
			0,
		},
		{
			100.49,
			100,
		},
	}
	for _, tc := range ts {
		i := int16(math.Round(tc.Float))
		require.Equal(t, tc.Int, i)
	}
}

func TestResize(t *testing.T) {
	image := make([][]YCbCrPixel, 12)
	for i := range image {
		image[i] = make([]YCbCrPixel, 10)
	}

	ts := []struct {
		Name    string
		Image   [][]YCbCrPixel
		X       int
		Y       int
		Changed bool
	}{
		{
			"NegativeNewValues", image, -123, -10, false,
		},
		{
			"EmptyImage", [][]YCbCrPixel{}, 10, 10, false,
		},
		{
			"ResizeSmallerSmaller", image, 8, 8, true,
		},
		{
			"ResizeSmallerBigger", image, 8, 16, true,
		},
		{
			"ResizeBiggerSmaller", image, 16, 8, true,
		},
		{
			"ResizeBiggerBigger", image, 16, 16, true,
		},
	}

	for _, tc := range ts {
		t.Run(tc.Name, func(t *testing.T) {
			rez, err := resize(tc.Image, tc.X, tc.Y)
			if !tc.Changed && err == nil {
				t.Fail()
			}
			if tc.Changed {
				require.EqualValues(t, tc.X, len(rez[0]))
				require.EqualValues(t, tc.Y, len(rez))
			} else {
				require.EqualValues(t, len(tc.Image), len(rez))
			}
		})
	}
}
