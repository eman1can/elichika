package live

import (
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_live"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

func surrender(ctx *gin.Context) {
	// there is no request body
	session := ctx.MustGet("session").(*userdata.Session)

	lpDiff := user_live.SurrenderLive(session)

	common.JsonResponse(ctx, &response.SurrenderLiveResponse{
		LpDiff:        lpDiff,
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	server.AddHandler("/", "POST", "/live/surrender", surrender)
}
