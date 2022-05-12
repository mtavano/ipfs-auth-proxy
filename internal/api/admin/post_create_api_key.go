package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mtavano/ipfs-auth-proxy/internal/database"
	"github.com/pkg/errors"
)

type postCreateApiKeyResponse struct {
	Key             string `json:"key"`
	Enabled         bool   `json:"enabled"`
	Requests        uint64 `json:"requests"`
	BytesTransfered int64  `json:"bytesTransfered"`
}

func (h *Handler) PostCreateApiKey(c *gin.Context) (interface{}, int, error) {
	value, exist := c.Get("admin")
	if !exist {
		return nil, http.StatusUnauthorized, nil
	}

	admin := value.(*database.Admin)

	key, err := h.ctx.IDGenerator.Generate()
	if err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, "admin: Handler.PostCreateApiKey IDGenerator.Generate error")
	}

	apiKey := &database.ApiKey{
		Key:     key,
		Enabled: true,
	}
	h.ctx.Storage.CreateApiKey(admin, apiKey)

	return &postCreateApiKeyResponse{
		Key:     apiKey.Key,
		Enabled: apiKey.Enabled,
	}, http.StatusCreated, nil
}
