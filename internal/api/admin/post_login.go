package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mtavano/ipfs-auth-proxy/internal/database"
	"github.com/pkg/errors"
)

type postLoginRequest struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type postLoginResponse struct {
	Token string `json:"token"`
}

func (h *Handler) PostLogin(c *gin.Context) (interface{}, int, error) {
	var req postLoginRequest
	err := c.BindJSON(&req)
	if err != nil {
		return nil, http.StatusBadRequest, errPayloadMismatch
	}

	admin, err := h.ctx.Storage.GetAdmin(req.User, req.Password)
	if errors.Is(err, database.ErrNotFound) {
		return nil, http.StatusNotFound, nil
	}

	token, err := h.ctx.IDGenerator.Generate()
	if err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, "admin: Handler.PostLogin IDGenerator.Generate error")
	}

	h.ctx.Storage.PutAdminToken(token, admin)

	return &postLoginResponse{Token: token}, http.StatusCreated, nil
}
