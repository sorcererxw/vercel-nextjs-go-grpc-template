package api

import (
	"net/http"

	"github.com/sorcererxw/vercel-nextjs-go-grpc-template/internal/app"
)

func Handle(w http.ResponseWriter, r *http.Request) { app.Handle(w, r) }
