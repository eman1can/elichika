package agnostic

import (
	"net/http"
	"os"
	"path/filepath"

	"elichika/internal/config"
	"elichika/internal/gamedata"
	"elichika/internal/utils"
	"elichika/internal/webui/request"

	"elichika/internal/server"

	"github.com/gin-gonic/gin"
)

// getVoiceSound handles GET /webui/item/sound?id=N
// It converts on first request then streams from cache.
func getVoiceSound(ctx *gin.Context) {
	req := request.WebUIItemSoundRequest{}
	if err := ctx.ShouldBindQuery(&req); err != nil || req.VoiceId == 0 {
		ctx.Status(http.StatusBadRequest)
		return
	}

	if voice, ok := gamedata.Instance.NaviVoice[req.VoiceId]; ok {
		wavPath, err := ConvertVoiceToWAV(voice.SheetName)
		if err != nil {
			ctx.Status(http.StatusNoContent)
			ctx.Error(err)
			return
		}

		ctx.Header("Content-Type", "audio/wav")
		ctx.File(wavPath)
	} else {
		ctx.Status(http.StatusNotFound)
	}
}

func init() {
	server.AddHandler("/webui", "GET", "/item/sound", getVoiceSound)
	err := os.MkdirAll(filepath.Join(config.StaticDataPath, "sounds", "wav"), os.ModePerm)
	utils.CheckErr(err)
}
