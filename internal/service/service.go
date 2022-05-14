package service

import (
	"github.com/sorcererxw/vercel-nextjs-go-grpc-template/internal/dao"
)

type Service struct {
	dao *dao.Dao
}

func New(d *dao.Dao) *Service {
	return &Service{dao: d}
}
