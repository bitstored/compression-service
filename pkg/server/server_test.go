package server

import (
	"context"
	"reflect"
	"testing"

	"github.com/bitstored/compression-service/pb"
	"github.com/bitstored/compression-service/pkg/service"
)

func TestNewServer(t *testing.T) {
	type args struct {
		s *service.Service
	}
	tests := []struct {
		name string
		args args
		want *Server
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServer(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_CompressImage(t *testing.T) {
	type fields struct {
		Service *service.Service
	}
	type args struct {
		ctx context.Context
		in  *pb.CompressImageRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.CompressImageResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Service: tt.fields.Service,
			}
			got, err := s.CompressImage(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.CompressImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.CompressImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_DecompressImage(t *testing.T) {
	type fields struct {
		Service *service.Service
	}
	type args struct {
		ctx context.Context
		in  *pb.DecompressImageRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.DecompressImageResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Service: tt.fields.Service,
			}
			got, err := s.DecompressImage(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.DecompressImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.DecompressImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_CompressText(t *testing.T) {
	type fields struct {
		Service *service.Service
	}
	type args struct {
		ctx context.Context
		in  *pb.CompressTextRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.CompressTextResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Service: tt.fields.Service,
			}
			got, err := s.CompressText(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.CompressText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.CompressText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_DecompressText(t *testing.T) {
	type fields struct {
		Service *service.Service
	}
	type args struct {
		ctx context.Context
		in  *pb.DecompressTextRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.DecompressTextResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Service: tt.fields.Service,
			}
			got, err := s.DecompressText(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.DecompressText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.DecompressText() = %v, want %v", got, tt.want)
			}
		})
	}
}
