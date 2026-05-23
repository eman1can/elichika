package live

import (
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_live"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

func fetchLiveMusicSelect(ctx *gin.Context) {
	// ther is no request body
	session := ctx.MustGet("session").(*userdata.Session)

	resp := user_live.FetchLiveMusicSelect(session)

	common.JsonResponse(ctx, &resp)
}

func init() {
	server.AddHandler("/", "POST", "/live/fetchLiveMusicSelect", fetchLiveMusicSelect)
}
