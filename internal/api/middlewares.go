package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mtavano/ipfs-auth-proxy/internal/app"
	"github.com/mtavano/ipfs-auth-proxy/internal/database"
)

const (
	adminTokenHeader = "x-admin-token"
	apiKeyHeader     = "x-api-key"
)

func adminAuthMiddleware(ctx *app.Context) func(*gin.Context) {
	return func(c *gin.Context) {
		token := c.GetHeader(adminTokenHeader)
		admin, err := ctx.Storage.GetAdminByToken(token)
		if errors.Is(err, database.ErrNotFound) {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		c.Set("admin", admin)
	}
}

func apiKeyAuthMiddleware(ctx *app.Context) func(*gin.Context) {
	return func(c *gin.Context) {
		key := c.GetHeader(apiKeyHeader)
		apiKey, err := ctx.Storage.GetApiKey(key)
		if errors.Is(err, database.ErrNotFound) {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		if !apiKey.Enabled {
			c.JSON(http.StatusForbidden, nil)
			c.Abort()
			return
		}

		c.Set("apikey", apiKey)
	}
}

func useCors(r gin.IRouter) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", adminTokenHeader, apiKeyHeader},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
}
