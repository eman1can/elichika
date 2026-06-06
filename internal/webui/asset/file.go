package asset

import (
	"crypto/md5"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendFile(ctx *gin.Context, data []byte, mime string) {
	currentEtag := fmt.Sprintf("%x", md5.Sum(data))
	clientEtag := ctx.GetHeader("If-None-Match")
	if clientEtag != "" && clientEtag == currentEtag {
		ctx.Header("Etag", currentEtag)
		ctx.Status(http.StatusNotModified)
	} else {
		ctx.Header("Etag", currentEtag)
		ctx.Data(http.StatusOK, mime, data)
	}
}
