package server

import (
	"bytes"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"image"
	"image/png"

	"github.com/bitstored/compression-service/pb"
	"github.com/bitstored/compression-service/pkg/service"
)

type Server struct {
	Service *service.Service
}

func NewServer(s *service.Service) *Server {
	return &Server{s}
}

func (s *Server) CompressImage(ctx context.Context, in *pb.CompressImageRequest) (*pb.CompressImageResponse, error) {
	reader := bytes.NewReader(in.GetImage())

	image, format, err := image.Decode(reader)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "image can't be decoded")
	}

	level := compressionLevel(in.GetLevel())
	imgType := in.GetType()
	bytes, err := s.Service.CompressImage(ctx, image, level, imgType)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "image can't be compressed")
	}

	response := &pb.CompressImageResponse{
		ResponseCode:    200,
		ResponseMessage: "Image compressed successfully",
		Image:           bytes,
		Format:          format,
	}
	return response, nil
}

func (s *Server) DecompressImage(ctx context.Context, in *pb.DecompressImageRequest) (*pb.DecompressImageResponse, error) {

	level := compressionLevel(in.GetLevel())
	imgType := in.GetType()

	image, err := s.Service.DecompressImage(ctx, in.GetImage(), level, imgType)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "image can't be decompressed")
	}

	writer := new(bytes.Buffer)
	png.Encode(writer, image)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "image can't be decoded")
	}

	bytes := writer.Bytes()
	response := &pb.DecompressImageResponse{
		ResponseCode:    200,
		ResponseMessage: "Image decompressed successfully",
		Image:           bytes,
	}
	return response, nil
}

func (s *Server) CompressText(context.Context, *pb.CompressTextRequest) (*pb.CompressTextResponse, error) {
	return nil, nil
}
func (s *Server) DecompressText(context.Context, *pb.DecompressTextRequest) (*pb.DecompressTextResponse, error) {
	return nil, nil
}
