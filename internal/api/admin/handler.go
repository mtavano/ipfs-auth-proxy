package admin

import (
	"errors"

	"github.com/mtavano/ipfs-auth-proxy/internal/app"
)

var (
	errPayloadMismatch = errors.New("payload mismatch")
)

type Handler struct {
	ctx *app.Context
}

func NewHandler(ctx *app.Context) *Handler {
	return &Handler{
		ctx: ctx,
	}
}
