package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mtavano/ipfs-auth-proxy/internal/database"
	"github.com/pkg/errors"
)

type patchAdminKeyRequest struct {
	Key     string `json:"key" binding:"required"`
	Enabled bool   `json:"enabled"`
}

type patchAdminKeyResponse struct {
	Key             string `json:"key"`
	Enabled         bool   `json:"enabled"`
	Requests        uint64 `json:"requests"`
	BytesTransfered int64  `json:"bytesTransfered"`
}

func (h *Handler) PatchApiKey(c *gin.Context) (interface{}, int, error) {
	value, exist := c.Get("admin")
	if !exist {
		return nil, http.StatusUnauthorized, nil
	}
	admin := value.(*database.Admin)

	var req patchAdminKeyRequest
	err := c.BindJSON(&req)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	apiKey, err := h.ctx.Storage.GetApiKey(req.Key)
	if errors.Is(err, database.ErrNotFound) {
		return nil, http.StatusNotFound, nil
	}

	apiKey.Enabled = req.Enabled
	h.ctx.Storage.PutApiKey(apiKey)

	switch true {
	case apiKey.Enabled && !req.Enabled:
		admin.DisabledApiKeys += 1
		break

	case !apiKey.Enabled && req.Enabled:
		admin.DisabledApiKeys -= 1
	}

	h.ctx.Storage.PutAdmin(admin)

	return &patchAdminKeyResponse{
		Key:             apiKey.Key,
		Enabled:         apiKey.Enabled,
		Requests:        apiKey.Requests,
		BytesTransfered: apiKey.BytesTransfered,
	}, http.StatusOK, nil
}
