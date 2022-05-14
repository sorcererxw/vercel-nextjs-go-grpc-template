package app

import (
	httppkg "net/http"

	"github.com/sorcererxw/vercel-nextjs-go-grpc-template/internal/config"
	"github.com/sorcererxw/vercel-nextjs-go-grpc-template/internal/dao"
	"github.com/sorcererxw/vercel-nextjs-go-grpc-template/internal/grpc"
	"github.com/sorcererxw/vercel-nextjs-go-grpc-template/internal/http"
	"github.com/sorcererxw/vercel-nextjs-go-grpc-template/internal/service"
)

var handler httppkg.HandlerFunc

func init() {
	conf, err := config.Parse()
	if err != nil {
		panic(err)
	}
	d, err := dao.New(conf)
	if err != nil {
		panic(err)
	}
	svc := service.New(d)
	httpSrv := http.New(svc)
	grpcSrc := grpc.New(svc)

	handler = func(w httppkg.ResponseWriter, r *httppkg.Request) {
		switch r.Header.Get("Content-Type") {
		case "application/grpc-web", "application/grpc-web-text":
			grpcSrc.ServeHTTP(w, r)
		default:
			httpSrv.ServeHTTP(w, r)
		}
	}
}

func Handle(w httppkg.ResponseWriter, r *httppkg.Request) {
	handler.ServeHTTP(w, r)
}
