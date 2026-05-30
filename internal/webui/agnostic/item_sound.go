package agnostic

import (
	"net/http"

	"elichika/internal/webui/request"

	"elichika/internal/server"

	"github.com/gin-gonic/gin"
)

// getVoiceSound handles GET /webui/item/sound?id=N&route=N&value=N
// It converts on first request then streams from cache.
func getVoiceSound(ctx *gin.Context) {
	req := request.WebUIItemSoundRequest{}
	if err := ctx.ShouldBindQuery(&req); err != nil || req.VoiceId == 0 {
		ctx.Status(http.StatusBadRequest)
		return
	}

	sheetName := NaviVoiceSheetName(req.ReleaseRoute, req.ReleaseValue, req.VoiceId)
	if sheetName == "" {
		ctx.Status(http.StatusNotFound)
		return
	}

	wavPath, err := ConvertVoiceToWAV(sheetName)
	if err != nil {
		ctx.Status(http.StatusNoContent)
		ctx.Error(err)
		return
	}

	ctx.Header("Content-Type", "audio/wav")
	ctx.File(wavPath)
}

func init() {
	server.AddHandler("/webui", "GET", "/item/sound", getVoiceSound)
}
