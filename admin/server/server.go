package server

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-reminder-bot/pkg/config"
)

type Server interface {
	Start(ctx context.Context)
	Stop(ctx context.Context)
}
type server struct {
	handler     Handler
	authHandler gin.HandlerFunc
}

func (s server) Stop(ctx context.Context) {
	//TODO implement me
}

func NewServer(handler Handler) Server {
	return &server{handler: handler}
}

func (s server) Start(ctx context.Context) {
	h := s.handler
	g := gin.Default()
	if config.DevelopmentMode == false {
		g.Static("/admin", "./admin/reminder-admin/dist")
	}
	g.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: false,
		MaxAge:           86400,
	}))
	r := g.Group("/api/v1")

	userG := r.Group("/users")
	{
		userG.GET("", h.findUsers)
		userG.GET("/:id", h.getOneUser)
		userG.POST("", h.createUser)
		userG.PUT("/:id", h.updateUser)
		userG.DELETE("/:id", h.deleteUser)
	}

	g.Run()
}
