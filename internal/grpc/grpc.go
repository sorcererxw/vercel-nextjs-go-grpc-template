package grpc

import (
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"

	v1 "github.com/sorcererxw/vercel-nextjs-go-grpc-template/idl_gen/go/api/v1"
	"github.com/sorcererxw/vercel-nextjs-go-grpc-template/internal/service"
)

type Handler http.Handler

func New(svc *service.Service) Handler {
	srv := grpc.NewServer()
	v1.RegisterAPIServer(srv, &v1Impl{svc: svc})

	return grpcweb.WrapServer(srv)
}
