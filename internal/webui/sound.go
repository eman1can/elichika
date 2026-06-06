package webui

import (
	"net/http"
	"os"
	"path/filepath"

	"elichika/internal/config"
	"elichika/internal/server"
	"elichika/internal/utils"
	"elichika/internal/webui/asset"

	"github.com/gin-gonic/gin"
)

type WebUIItemSoundRequest struct {
	SoundAssetPath string `form:"sound_asset_path"`
	Language       string `form:"l"`
}

// getVoiceSound handles GET /webui/item/sound?id=N
// It converts on first request then streams from cache.
func getVoiceSound(ctx *gin.Context) {
	var req WebUIItemSoundRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	output, err := asset.ConvertVoiceToWAV(ctx, req.SoundAssetPath)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	asset.SendFile(ctx, output, "audio/wav")
}

func init() {
	server.AddHandler("/webui", "GET", "/sound", getVoiceSound)
	err := os.MkdirAll(filepath.Join(config.StaticDataPath, "sounds", "wav"), os.ModePerm)
	utils.CheckErr(err)
}
