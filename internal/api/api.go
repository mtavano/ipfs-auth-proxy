package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mtavano/ipfs-auth-proxy/config"
	"github.com/mtavano/ipfs-auth-proxy/internal/api/admin"
	v1 "github.com/mtavano/ipfs-auth-proxy/internal/api/v1"
	"github.com/mtavano/ipfs-auth-proxy/internal/app"
)

type Api struct {
	port   string
	router *gin.Engine
	ctx    *app.Context
}

func New(ctx *app.Context, c *config.Config) *Api {
	return &Api{
		port:   c.Port,
		router: gin.Default(),
		ctx:    ctx,
	}
}
func (api *Api) Run() {
	useCors(api.router)
	adminAuth := adminAuthMiddleware(api.ctx)
	apiKeyAuth := apiKeyAuthMiddleware(api.ctx)

	// Admin endpoints
	adminHandler := admin.NewHandler(api.ctx)
	api.router.POST("/admin/login", api.handleFunc(adminHandler.PostLogin))
	api.router.POST("/admin/apikeys", adminAuth, api.handleFunc(adminHandler.PostCreateApiKey))
	api.router.PATCH("/admin/apikeys", adminAuth, api.handleFunc(adminHandler.PatchApiKey))
	api.router.GET("/admin/apikeys", adminAuth, api.handleFunc(adminHandler.GetApiKeys))

	// v1 endpoints
	v1Handler := v1.NewHandler(api.ctx)
	api.router.GET("/v1/files/:cid", apiKeyAuth, v1Handler.GetFile)

	api.router.Run(fmt.Sprintf(":%s", api.port))
}

type handler func(*gin.Context) (interface{}, int, error)

func (s *Api) handleFunc(fn handler) func(*gin.Context) {
	return func(c *gin.Context) {
		payload, statusCode, err := fn(c)

		if err != nil {
			c.JSON(statusCode, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(statusCode, payload)
	}
}
