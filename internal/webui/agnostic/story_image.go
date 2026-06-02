package agnostic

import (
	"crypto/md5"
	"fmt"
	"net/http"

	"elichika/internal/gamedata"
	"elichika/internal/server"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

type WebUIStoryImageRequest struct {
	Id int32 `form:"id"`
}

func getStoryImage(ctx *gin.Context) {
	req := WebUIStoryImageRequest{}
	err := ctx.ShouldBindQuery(&req)
	utils.CheckErr(err)

	if story, ok := gamedata.Instance.StoryMainChapter[req.Id]; ok {
		if story.ThumbnailAssetPath == "" {
			ctx.Status(http.StatusNotFound)
			return
		}

		output, err := loadAssetImage(story.ThumbnailAssetPath)
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
		return
	}

	ctx.Status(http.StatusBadRequest)
}

func init() {
	server.AddHandler("/webui", "GET", "/story/image", getStoryImage)
}
