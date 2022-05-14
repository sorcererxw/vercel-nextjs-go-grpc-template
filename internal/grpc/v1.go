package grpc

import (
	"context"

	v1 "github.com/sorcererxw/vercel-nextjs-go-grpc-template/idl_gen/go/api/v1"
	"github.com/sorcererxw/vercel-nextjs-go-grpc-template/internal/service"
)

type v1Impl struct {
	v1.UnsafeAPIServer

	svc *service.Service
}

func (v *v1Impl) Ping(ctx context.Context, request *v1.PingRequest) (*v1.PingResponse, error) {
	return &v1.PingResponse{
		Message: "pong",
	}, nil
}
