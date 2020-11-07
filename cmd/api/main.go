package main

import (
	"api/config"
	lg "api/logger"
	"api/pkg/blog/repo"
	"api/pkg/blog/service"
	"api/pkg/presentation"
	v1 "api/pkg/presentation/v1"
	"context"
)

func main() {
	ctx := context.Background()

	config, e := config.GetDefault(ctx)
	if e != nil {
		lg.Fatal(ctx, e)
	}

	// blog
	repo, e := repo.New(ctx, config.MariaDB)
	if e != nil {
		lg.Fatal(ctx, e)
	}

	srv := service.New(repo)
	if e != nil {
		lg.Fatal(ctx, e)
	}

	// router
	r := presentation.NewRouter()

	// handlers
	base := presentation.NewBaseHandler(ctx, r)
	v1, e := v1.NewHandler(config, srv, r)
	if e != nil {
		lg.Fatal(ctx, e)
	}

	presentation.New(ctx, &config.Service, r, []presentation.Handler{base, v1})

	// port := ":" + config.Service.Port
	// lg.Infof(ctx, "service running on port: %d", port)
	// log.Fatal(
	// 	http.ListenAndServe(port, presentation.Router.Get()))
}
