package mission

import (
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_mission"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

func fetchMission(ctx *gin.Context) {
	// there is no request body
	session := ctx.MustGet("session").(*userdata.Session)

	resp := user_mission.FetchMission(session)

	common.JsonResponse(ctx, &resp)
}

func init() {
	server.AddHandler("/", "POST", "/mission/fetchMission", fetchMission)
}
