package user

import (
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_status"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func recoverLpSubscription(ctx *gin.Context) {
	// there is no request body
	session := ctx.MustGet("session").(*userdata.Session)

	user_status.AddUserLp(session, session.Gamedata.UserRank[session.UserStatus.Rank].MaxLp)
	session.UserStatus.LivePointSubscriptionRecoveryDailyCount = 1 // 1 mean used
	session.UserStatus.LivePointSubscriptionRecoveryDailyResetAt = utils.BeginOfNextDay(session.Time).Unix()

	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	server.AddHandler("/", "POST", "/user/recoverLpSubscription", recoverLpSubscription)
}
