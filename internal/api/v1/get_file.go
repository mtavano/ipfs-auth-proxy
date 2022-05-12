package v1

import (
	"bytes"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipfs/interface-go-ipfs-core/path"
	"github.com/mtavano/ipfs-auth-proxy/internal/database"
)

func (h *Handler) GetFile(c *gin.Context) {
	value, exist := c.Get("apikey")
	if !exist {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	apiKey := value.(*database.ApiKey)

	cid := c.Param("cid")

	fileReader, err := h.ctx.BlockGetter.Get(context.Background(), path.New(cid))
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	fileBuff := new(bytes.Buffer)
	fileSize, err := fileBuff.ReadFrom(fileReader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	apiKey.Requests += uint64(1)
	apiKey.BytesTransfered += fileSize
	h.ctx.Storage.PutApiKey(apiKey)

	c.Writer.Write(fileBuff.Bytes())
	c.Writer.WriteHeader(http.StatusOK)
	return
}
