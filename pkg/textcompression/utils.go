package textcompression

import (
	"compress/flate"
	"github.com/bitstored/compression-service/pb"
)

func compressionLevel(l pb.CompressionLevel) int {
	switch l {
	case pb.CompressionLevel_DefaultCompression:
		return flate.DefaultCompression
	case pb.CompressionLevel_NoCompression:
		return flate.NoCompression
	case pb.CompressionLevel_BestSpeed:
		return flate.BestSpeed
	case pb.CompressionLevel_BestCompression:
		return flate.BestCompression
	case pb.CompressionLevel_HuffmanOnly:
		return flate.HuffmanOnly
	default:
		return flate.DefaultCompression
	}
}
