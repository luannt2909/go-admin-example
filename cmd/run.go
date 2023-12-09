package cmd

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-admin/di"
	"go-admin/server"
	"go.uber.org/fx"
	"os"
)

func Execute() error {
	app := fx.New(
		di.Module,
		fx.Invoke(startAdminServer),
	)
	app.Run()
	return nil
}

func serveStaticFile(g *gin.Engine) {
	g.Static("/admin", "./webs/dist")
}

func startAdminServer(lc fx.Lifecycle, router server.Router) {
	g := gin.Default()
	serveStaticFile(g)
	router.Register(g)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			port := os.Getenv("PORT")
			if port == "" {
				port = "2303"
			}
			go g.Run(":" + port)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
