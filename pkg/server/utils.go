package server

import (
	"github.com/bitstored/compression-service/pb"
	"image/png"
)

func compressionLevel(l pb.CompressionLevel) png.CompressionLevel {
	switch l {
	case pb.CompressionLevel_DefaultCompression:
		return png.DefaultCompression
	case pb.CompressionLevel_NoCompression:
		return png.NoCompression
	case pb.CompressionLevel_BestSpeed:
		return png.BestSpeed
	case pb.CompressionLevel_BestCompression:
		return png.BestCompression
	default:
		return png.DefaultCompression
	}
}
