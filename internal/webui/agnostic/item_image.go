package agnostic

import (
	"crypto/md5"
	"fmt"
	"net/http"

	"elichika/internal/server"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

type WebUIItemImageRequest struct {
	ContentType int32 `form:"type"`
	ContentId   int32 `form:"id"`
}

func getItemImage(ctx *gin.Context) {
	req := WebUIItemImageRequest{}
	err := ctx.ShouldBindQuery(&req)
	utils.CheckErr(err)

	assetPath := ItemsByItemId[req.ContentType][req.ContentId].AssetPath
	if assetPath == "" {
		ctx.Status(http.StatusNotFound)
		return
	}

	output, err := loadAssetImage(assetPath)
	utils.CheckErr(err)

	currentEtag := fmt.Sprintf("%x", md5.Sum(output))
	clientEtag := ctx.GetHeader("If-None-Match")
	if clientEtag != "" && clientEtag == currentEtag {
		ctx.Header("Etag", currentEtag)
		ctx.Status(http.StatusNotModified)
	} else {
		ctx.Header("Etag", currentEtag)
		ctx.Data(http.StatusOK, "image/png", output)
	}
}

func init() {
	server.AddHandler("/webui", "GET", "/item/image", getItemImage)
}
