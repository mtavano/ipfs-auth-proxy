package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type getApiKeyResponse struct {
	ApiKeys []*apiKey `json:"apiKeys"`
}

type apiKey struct {
	Key             string `json:"key"`
	Enabled         bool   `json:"enabled"`
	Requests        uint64 `json:"requests"`
	BytesTransfered int64  `json:"bytesTransfered"`
}

func (h *Handler) GetApiKeys(c *gin.Context) (interface{}, int, error) {
	apiKeys, err := h.ctx.Storage.GetApiKeys()
	if err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, "admin Handler.GetApiKeys Storage.GetApiKeys error")
	}

	res := &getApiKeyResponse{
		ApiKeys: make([]*apiKey, 0),
	}

	for _, ak := range apiKeys {
		res.ApiKeys = append(res.ApiKeys, &apiKey{
			Key:             ak.Key,
			Enabled:         ak.Enabled,
			Requests:        ak.Requests,
			BytesTransfered: ak.BytesTransfered,
		})
	}

	return res, http.StatusOK, nil
}
