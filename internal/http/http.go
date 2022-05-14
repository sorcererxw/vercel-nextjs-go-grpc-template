package http

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sorcererxw/vercel-nextjs-go-grpc-template/internal/service"
)

type Handler http.Handler

func New(svc *service.Service) Handler {
	e := echo.New()

	e.GET("/api/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	return e
}
