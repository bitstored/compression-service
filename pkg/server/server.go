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

	if in.GetImage() == nil || len(in.GetImage()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "image can't be empty")
	}

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

	if in.GetImage() == nil || len(in.GetImage()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "image can't be empty")
	}

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

func (s *Server) CompressText(ctx context.Context, in *pb.CompressTextRequest) (*pb.CompressTextResponse, error) {
	level := in.GetLevel()

	if in.GetText() == nil || len(in.GetText()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "text can't be empty")
	}
	out, err := s.Service.CompressText(ctx, in.GetText(), level)
	if err != nil {
		return nil, err
	}

	rsp := pb.CompressTextResponse{
		ResponseCode:    200,
		ResponseMessage: "Text compressed succesfully",
		Text:            out,
		Level:           level,
	}

	return &rsp, nil
}
func (s *Server) DecompressText(ctx context.Context, in *pb.DecompressTextRequest) (*pb.DecompressTextResponse, error) {

	if in.GetText() == nil || len(in.GetText()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "text can't be empty")
	}
	out, err := s.Service.DecompressText(ctx, in.GetText())
	if err != nil {
		return nil, err
	}

	rsp := pb.DecompressTextResponse{
		ResponseCode:    200,
		ResponseMessage: "Text decompressed succesfully",
		Text:            out,
	}

	return &rsp, nil
}
