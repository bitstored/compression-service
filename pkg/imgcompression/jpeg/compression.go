package jpeg

import (
	"math"

	"github.com/bitstored/compression-service/pkg/errors"
)

/* multiply multiplies the RGBPixel vector by an 3x3 matrix
 * @returns the resulted YCbCr pixel vector
 */
func (p RGBPixel) multiply(matrix [][]float64) YCbCrPixel {
	out := YCbCrPixel{}
	out.Y = maxInt16(0, minInt16(math.MaxUint8, mul(matrix[0][0], p.R)+mul(matrix[0][1], p.G)+mul(matrix[0][2], p.B)))
	out.Cb = maxInt16(0, minInt16(math.MaxUint8, transformationValue+mul(matrix[1][0], p.R)+mul(matrix[1][1], p.G)+mul(matrix[1][2], p.B)))
	out.Cr = maxInt16(0, minInt16(math.MaxUint8, transformationValue+mul(matrix[2][0], p.R)+mul(matrix[2][1], p.G)+mul(matrix[2][2], p.B)))
	return out
}

/* toYCbCr transforms a RGB image to an YCbCr image
 * @returns the result of the conversion
 */
func toYCbCr(in [][]RGBPixel) ([][]YCbCrPixel, *errors.Err) {

	if in == nil || len(in) == 0 || len(in[0]) == 0 {
		return nil, errors.NewError(errors.KindEmptyImage, "image can not be converted to YCbCr, image is empty")
	}

	out := make([][]YCbCrPixel, len(in))
	for i := range in {
		out[i] = make([]YCbCrPixel, len(in[i]))
		for j := range in[i] {
			out[i][j] = in[i][j].multiply(RGB2YcbCrMatrix)
		}
	}

	return out, nil
}

// toSigned converts an uint8 to uint8
func toSigned(in [][]YCbCrPixel) [][]YCbCrPixel {
	for i := range in {
		for j := range in[i] {
			in[i][j].Y -= transformationValue
		}
	}
	return in
}

/* extract8x8Tile extracts an 8x8 submatrix from the input
 * matrix up-left corner is indiocated by x an d y
 * @returns the matrix if this can be extracted, ohtherwise error
 */
func extract8x8Tile(in [][]YCbCrPixel, x int, y int) ([][]YCbCrPixel, *errors.Err) {

	if len(in) < x+blockLen || x < 0 {
		return nil, errors.NewError(errors.KindOverflow, "unable to extract a 8x8 submatrix, x index is overflowing")
	}

	if len(in[0]) < y+blockLen || y < 0 {
		return nil, errors.NewError(errors.KindOverflow, "unable to extract a 8x8 submatrix, y index is overflowing")
	}

	out := make([][]YCbCrPixel, 8)
	for i := 0; i < blockLen; i++ {
		out[i] = in[x+i][y : y+blockLen]
	}

	return out, nil
}

/* funcDCT applies discrete cosine transform to the 8x8 tile
 * @returns the result of the transformation
 */
func funcDCT(tile [][]YCbCrPixel) [][]YCbCrPixel {

	out := make([][]YCbCrPixel, blockLen)
	for i := range tile[0] {
		out[i] = make([]YCbCrPixel, blockLen)
	}

	for i := 0; i < blockLen; i++ {
		for j := 0; j < blockLen; j++ {
			// compute ci for FDCT
			var ci float64

			if i == 0 {
				ci = math.Sqrt(1.0 / float64(blockLen))
			} else {
				ci = math.Sqrt(2.0 / float64(blockLen))
			}
			// compute cj for FDCT
			var cj float64

			if j == 0 {
				cj = math.Sqrt(1.0 / float64(blockLen))
			} else {
				cj = math.Sqrt(2.0 / float64(blockLen))
			}
			var auxY float64
			var auxCb float64
			var auxCr float64
			for x := 0; x < blockLen; x++ {
				for y := 0; y < blockLen; y++ {
					cosYi := math.Cos((2*float64(y) + 1) * math.Pi * float64(i) / 16)
					cosXj := math.Cos((2*float64(x) + 1) * math.Pi * float64(j) / 16)
					auxY += float64(tile[x][y].Y) * cosXj * cosYi
					auxCb += float64(tile[x][y].Cb) * cosXj * cosYi
					auxCr += float64(tile[x][y].Cr) * cosXj * cosYi
				}
			}

			out[i][j].Y = int16(math.Round(auxY * ci * cj))
			out[i][j].Cr = int16(math.Round(auxCr * ci * cj))
			out[i][j].Cb = int16(math.Round(auxCb * ci * cj))
		}
	}

	return out
}

func quantize(tile [][]YCbCrPixel) ([][]YCbCrPixel, *errors.Err) {

	if len(tile) != blockLen || len(tile[0]) != blockLen {
		return nil, errors.NewError(errors.KindMalformedStructure, "tile can not be encodded")
	}

	for x := 0; x < blockLen; x++ {
		for y := 0; y < blockLen; y++ {
			val := float64(tile[x][y].Y) / float64(quantizationMatrix[x][y])
			val = math.Round(val)
			tile[x][y].Y = int16(val)

			val = float64(tile[x][y].Cr) / float64(chromianceMatrix[x][y])
			val = math.Round(val)
			tile[x][y].Cr = int16(val)

			val = float64(tile[x][y].Cb) / float64(chromianceMatrix[x][y])
			val = math.Round(val)
			tile[x][y].Cb = int16(val)
		}
	}

	return tile, nil
}

func encodeEntropy(tile [][]YCbCrPixel) (yArray []int8, cbArray []int8, crArray []int8, err *errors.Err) {

	if len(tile) != blockLen || len(tile[0]) != blockLen {
		return nil, nil, nil, errors.NewError(errors.KindMalformedStructure, "tile can not be encodded")
	}

	yArray = append(yArray, 0)   // number of values in Y
	cbArray = append(cbArray, 0) // number of values in Cb
	crArray = append(crArray, 0) // number of values in Cr

	var y = Pair{int8(tile[0][0].Y), 1}
	var cb = Pair{int8(tile[0][0].Cb), 1}
	var cr = Pair{int8(tile[0][0].Cr), 1}

	// zig zag plm
	for i := 1; i < blockLen; i++ {
		for j := 0; j <= i; j++ {
			var yAux, cbAux, crAux int8

			if i%2 == 1 {
				yAux = int8(tile[j][i-j].Y)
				cbAux = int8(tile[j][i-j].Cb)
				crAux = int8(tile[j][i-j].Cr)
			} else {
				yAux = int8(tile[i-j][j].Y)
				cbAux = int8(tile[i-j][j].Cb)
				crAux = int8(tile[i-j][j].Cr)
			}

			if yAux != y.Value {
				yArray = append(yArray, y.Value, y.Freq)
				y.Value = 0
				y.Freq = 0
			}

			if cbAux != cb.Value {
				cbArray = append(cbArray, cb.Value, cb.Freq)
				cb.Value = 0
				cb.Freq = 0
			}

			if crAux != cr.Value {
				crArray = append(crArray, cr.Value, cr.Freq)
				cr.Value = 0
				cr.Freq = 0
			}

			y.Value = yAux
			y.Freq++
			cb.Value = cbAux
			cb.Freq++
			cr.Value = crAux
			cr.Freq++
		}
	}

	// Add last value if is != 0
	if y.Value != 0 || len(yArray) == 0 {
		yArray = append(yArray, y.Value, y.Freq)
	}

	if cr.Value != 0 || len(yArray) == 0 {
		crArray = append(crArray, cr.Value, cr.Freq)
	}

	if cb.Value != 0 || len(yArray) == 0 {
		cbArray = append(cbArray, cb.Value, cb.Freq)
	}

	// add the number of values to read for the tile
	yArray[0] = int8((len(yArray) - 1) / 2)
	cbArray[0] = int8((len(cbArray) - 1) / 2)
	crArray[0] = int8((len(crArray) - 1) / 2)

	return yArray, cbArray, crArray, nil
}

/* applyCosine applies cosine transform to each 8x8 tile
 * @returns the resulted matrix
 */
func applyCosineAndEncode(in [][]YCbCrPixel) []int8 {
	out := make([]int8, 0)
	out = append(out, int2Int8Array(len(in))...)
	out = append(out, int2Int8Array(len(in[0]))...)

	for x := 0; x < len(in); x += blockLen {
		for y := 0; y < len(in[0]); y += blockLen {
			tile, err := extract8x8Tile(in, x, y)

			if err != nil {
				panic(err.Error())
			}

			dctTile := funcDCT(tile)
			dctTile, err = quantize(dctTile)
			if err != nil {
				panic(err.Error())
			}

			y, cb, cr, err := encodeEntropy(dctTile)
			if err != nil {
				panic(err.Error())
			}

			out = append(out, y...)
			out = append(out, cb...)
			out = append(out, cr...)
		}
	}
	return out
}
