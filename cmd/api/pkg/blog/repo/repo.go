package repo

import (
	"context"
	"errors"

	"api/config"
	"api/logger"
	"api/pkg/blog"
	"api/pkg/blog/repo/mock"
	"api/pkg/blog/repo/mysql"
)

func New(ctx context.Context,
	conf config.DataSource) (r blog.Repo, e error) {

	switch conf.Engine {
	case "mock":
		return mock.New(ctx, conf)
	case "mysql":
		return mysql.New(ctx, conf)
	default:
		e = errors.New("unsupported or missing database engine")
		logger.Error(ctx, e)
		return
	}
}
