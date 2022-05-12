package v1

import "github.com/mtavano/ipfs-auth-proxy/internal/app"

type Handler struct {
	ctx *app.Context
}

func NewHandler(ctx *app.Context) *Handler {
	return &Handler{
		ctx: ctx,
	}
}
