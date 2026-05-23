package live

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_live"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func skip(ctx *gin.Context) {
	req := request.SkipLiveRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	resp := user_live.SkipLive(session, req)

	common.JsonResponse(ctx, &resp)
}

func init() {
	server.AddHandler("/", "POST", "/live/skip", skip)
}
