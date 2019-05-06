package imgcompression

type RGBPixel struct {
	R int16
	G int16
	B int16
}

type YCbCrPixel struct {
	Y  int16
	Cb int16
	Cr int16
}

type Pixel interface {
	multiply(matrix [][]float64) Pixel
}
