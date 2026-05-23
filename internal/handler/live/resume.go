package live

import (
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/time"
	"elichika/internal/subsystem/user_live"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func resume(ctx *gin.Context) {
	// there is no request body

	session := ctx.MustGet("session").(*userdata.Session)

	exist, live, startReq := user_live.LoadUserLive(session)
	utils.MustExist(exist)

	common.JsonResponse(ctx, &response.ResumeLiveResponse{
		Live:          live,
		PartnerUserId: startReq.PartnerUserId,
		IsAutoPlay:    startReq.IsAutoPlay,
		WeekdayState:  time.GetWeekdayState(session),
	})
}

func init() {
	server.AddHandler("/", "POST", "/live/resume", resume)
}
