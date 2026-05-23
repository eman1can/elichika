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

// response: RecoverDailyLiveMusicPlayableResponse
// alternative respnose: RecoverableExceptionResponse
func recoverDailyLiveMusicPlayable(ctx *gin.Context) {
	req := request.RecoverDailyLiveMusicPlayableRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	successResponse, failureResponse := user_live.RecoverDailyLiveMusicPlayable(session, req.LiveId)
	if successResponse != nil {
		common.JsonResponse(ctx, successResponse)
	} else {
		common.AlternativeJsonResponse(ctx, failureResponse)
	}
}

func init() {
	server.AddHandler("/", "POST", "/live/recoverDailyLiveMusicPlayable", recoverDailyLiveMusicPlayable)
}
