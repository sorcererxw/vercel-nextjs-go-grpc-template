package dao

import (
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	_ "github.com/go-sql-driver/mysql"

	"github.com/sorcererxw/vercel-nextjs-go-grpc-template/internal/config"
)

type Dao struct {
	redis *redis.Client
	db    *sqlx.DB
}

func New(conf config.Config) (d *Dao, err error) {
	d = new(Dao)

	{
		opt, err := redis.ParseURL(conf.Redis.DSN)
		if err != nil {
			return nil, errors.Wrap(err, "parse redis dsn")
		}

		d.redis = redis.NewClient(opt)
	}

	{
		d.db, err = sqlx.Open("mysql", conf.Mysql.DSN)
		if err != nil {
			return nil, errors.Wrap(err, "open mysql")
		}
	}

	return &Dao{}, nil
}
