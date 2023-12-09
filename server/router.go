package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router interface {
	Register(g gin.IRouter)
}
type router struct {
	handler Handler
	authMdw AuthMiddleware
}

func NewRouter(handler Handler, authMdw AuthMiddleware) Router {
	return &router{handler: handler, authMdw: authMdw}
}

func (s router) Register(g gin.IRouter) {
	h := s.handler
	g.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: false,
		MaxAge:           86400,
	}))
	r := g.Group("/api/v1")
	r.POST("/auth/authenticate", h.Authenticate)
	userG := r.Group("/users", s.authMdw)
	{
		userG.GET("", h.findUsers)
		userG.GET("/:id", h.getOneUser)
		userG.POST("", h.createUser)
		userG.PUT("/:id", h.updateUser)
		userG.DELETE("/:id", h.deleteUser)
	}
}
