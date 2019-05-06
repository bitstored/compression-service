package imgcompression

import (
	"math"

	"github.com/bitstored/compression-service/pkg/errors"
)

const (
	valuesStart = 1
)

func toUnsigned(in [][]YCbCrPixel) [][]YCbCrPixel {

	for i := range in {
		for j := range in[i] {
			in[i][j].Y += transformationValue
			in[i][j].Y = minInt16(255, in[i][j].Y)
			in[i][j].Y = minInt16(255, in[i][j].Y)

		}
	}
	return in
}

func (p YCbCrPixel) multiply(matrix [][]float64) RGBPixel {
	out := RGBPixel{}
	out.R = maxInt16(0, minInt16(math.MaxUint8, mul(matrix[0][0], p.Y)+mul(matrix[0][1], p.Cb-transformationValue)+mul(matrix[0][2], p.Cr-transformationValue)))
	out.G = maxInt16(0, minInt16(math.MaxUint8, mul(matrix[1][0], p.Y)+mul(matrix[1][1], p.Cb-transformationValue)+mul(matrix[1][2], p.Cr-transformationValue)))
	out.B = maxInt16(0, minInt16(math.MaxUint8, mul(matrix[2][0], p.Y)+mul(matrix[2][1], p.Cb-transformationValue)+mul(matrix[2][2], p.Cr-transformationValue)))
	return out
}

/* toRGB transforms a YCbCr image to an RGB image
 * @returns the result of the conversion
 */
func toRGB(in [][]YCbCrPixel) ([][]RGBPixel, *errors.Err) {

	if in == nil || len(in) == 0 || len(in[0]) == 0 {
		return nil, errors.NewError(errors.KindEmptyImage, "image can not be converted to YCbCr, image is empty")
	}
	out := make([][]RGBPixel, len(in))
	for i := range in {
		out[i] = make([]RGBPixel, len(in[i]))
		for j := range in[i] {
			out[i][j] = in[i][j].multiply(YcbCr2RGBMatrix)
		}
	}

	return out, nil
}

func decodeEntropyTile(slice []int8) ([][]YCbCrPixel, *errors.Err) {
	if slice == nil {
		return nil, errors.NewError(errors.KindMalformedStructure, "slice is empty")
	}
	tile := make([][]YCbCrPixel, blockLen)
	for i := 0; i < blockLen; i++ {
		tile[i] = make([]YCbCrPixel, blockLen)
	}
	totCnt := 2 * int(slice[0])
	indx := int(valuesStart)
	cnt := int(slice[indx+1])
	val := int16(slice[indx])
	start := valuesStart
	for i := 0; i < blockLen; i++ {
		for j := 0; j <= i; j++ {
			if indx >= totCnt+start {
				goto decodeCb
			}
			if cnt == 0 {
				indx += 2
				if indx >= totCnt+start {
					goto decodeCb
				}
				cnt = int(slice[indx+1])
				val = int16(slice[indx])
			}
			if i%2 == 1 {
				tile[j][i-j].Y = val
			} else {
				tile[i-j][j].Y = val

			}
			cnt--
		}
	}
decodeCb:
	totCnt = 2 * int(slice[indx])
	indx += int(valuesStart)
	cnt = int(slice[indx+1])
	val = int16(slice[indx])
	start = indx
	for i := 0; i < blockLen; i++ {
		for j := 0; j <= i; j++ {
			if indx >= totCnt+start {
				goto decodeCr
			}
			if cnt == 0 {
				indx += 2
				if indx >= totCnt+start {
					goto decodeCr
				}
				cnt = int(slice[indx+1])
				val = int16(slice[indx])
			}
			if i%2 == 1 {
				tile[j][i-j].Cb = val
			} else {
				tile[i-j][j].Cb = val

			}
			cnt--
		}
	}
decodeCr:
	totCnt = 2 * int(slice[indx])
	indx += int(valuesStart)
	start = indx
	if indx >= totCnt+start {
		return tile, nil
	}
	cnt = int(slice[indx+1])
	val = int16(slice[indx])
	for i := 0; i < blockLen; i++ {
		for j := 0; j <= i; j++ {
			if indx >= totCnt+start {
				break
			}
			if cnt == 0 {
				indx += 2
				if indx >= totCnt+start {
					break
				}
				cnt = int(slice[indx+1])
				val = int16(slice[indx])
			}
			if i%2 == 1 {
				tile[j][i-j].Cr = val
			} else {
				tile[i-j][j].Cr = val
			}
			cnt--
		}
	}
	return tile, nil
}

func dequantize(tile [][]YCbCrPixel) ([][]YCbCrPixel, *errors.Err) {

	if len(tile) != blockLen || len(tile[0]) != blockLen {
		return nil, errors.NewError(errors.KindMalformedStructure, "tile can not be encodded")
	}

	for x := 0; x < blockLen; x++ {
		for y := 0; y < blockLen; y++ {
			val := float64(tile[x][y].Y) * float64(quantizationMatrix[x][y])
			val = math.Round(val)
			tile[x][y].Y = int16(val)

			val = float64(tile[x][y].Cr) * float64(chromianceMatrix[x][y])
			val = math.Round(val)
			tile[x][y].Cr = int16(val)

			val = float64(tile[x][y].Cb) * float64(chromianceMatrix[x][y])
			val = math.Round(val)
			tile[x][y].Cb = int16(val)
		}
	}
	return tile, nil
}

func funcICT(tile [][]YCbCrPixel) [][]YCbCrPixel {

	out := make([][]YCbCrPixel, blockLen)
	for i := range tile[0] {
		out[i] = make([]YCbCrPixel, blockLen)
	}

	for i := 0; i < blockLen; i++ {
		for j := 0; j < blockLen; j++ {

			var auxY float64
			var auxCb float64
			var auxCr float64
			for x := 0; x < blockLen; x++ {
				for y := 0; y < blockLen; y++ {
					// compute ci for FICT
					var ci float64
					if x == 0 {
						ci = math.Sqrt(1.0 / float64(blockLen))
					} else {
						ci = math.Sqrt(2.0 / float64(blockLen))
					}
					// compute cj for FICT
					var cj float64
					if y == 0 {
						cj = math.Sqrt(1.0 / float64(blockLen))
					} else {
						cj = math.Sqrt(2.0 / float64(blockLen))
					}

					cosXj := math.Cos((2*float64(j) + 1) * math.Pi * float64(x) / 16)
					cosYi := math.Cos((2*float64(i) + 1) * math.Pi * float64(y) / 16)

					auxY += ci * cj * float64(tile[x][y].Y) * cosXj * cosYi
					auxCb += ci * cj * float64(tile[x][y].Cb) * cosXj * cosYi
					auxCr += ci * cj * float64(tile[x][y].Cr) * cosXj * cosYi
				}
			}
			out[i][j].Y = int16(math.Round(auxY))
			out[i][j].Cr = int16(math.Round(auxCr))
			out[i][j].Cb = int16(math.Round(auxCb))
		}
	}

	return out
}

func decodeEntropyAndApplyCosine(data []int8) ([][]YCbCrPixel, *errors.Err) {
	X := int8Array2Int(data[0:4])
	Y := int8Array2Int(data[4:8])
	out := make([][]YCbCrPixel, X)
	for i := 0; i < X; i++ {
		out[i] = make([]YCbCrPixel, Y)
	}
	start := 8
	for i := 0; i < X; i += blockLen {
		for j := 0; j < Y; j += blockLen {
			// decode
			len := 2 * data[start]
			len += 2*data[start+int(len)+1] + 1
			len += 2*data[start+int(len)+1] + 1
			tile, err := decodeEntropyTile(data[start : start+int(len)+1])
			if err != nil {
				return nil, err
			}
			start += int(len) + 1
			// apply dequantization
			tile, err = dequantize(tile)
			if err != nil {
				return nil, err
			}
			//apply cosine
			tile = funcICT(tile)
			for x := 0; x < blockLen; x++ {
				for y := 0; y < blockLen; y++ {
					out[i+x][j+y] = tile[x][y]
				}
			}
		}
	}
	return out, nil
}

func decompressImage(data []int8) ([][]RGBPixel, *errors.Err) {
	image, err := decodeEntropyAndApplyCosine(data)
	if err != nil {
		return nil, err
	}

	x := len(data) - 8
	y := len(data) - 4
	image, err = resize(image, int8Array2Int(data[x:x+4]), int8Array2Int(data[y:y+4]))
	if err != nil {
		return nil, err
	}

	image = toUnsigned(image)

	out, err := toRGB(image)
	if err != nil {
		return nil, err
	}
	return out, nil
}
