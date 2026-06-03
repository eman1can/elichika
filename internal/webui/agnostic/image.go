package agnostic

import (
	"log"
	"net/http"

	"elichika/internal/server"

	"github.com/gin-gonic/gin"
)

type WebUIImageRequest struct {
	ImageAssetPath string `json:"image_asset_path" form:"image_asset_path" binding:"required"`
}

func getImage(ctx *gin.Context) {
	var req WebUIImageRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Print(err)
		return
	}

	log.Println("Received request for image asset path:", req.ImageAssetPath)
	output, err := loadAssetImage(req.ImageAssetPath)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Print(err)
		return
	}

	file(ctx, output, "image/png")
}

func init() {
	server.AddHandler("/webui", "GET", "/image", getImage)
}
