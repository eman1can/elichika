package info_trigger

import (
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_info_trigger"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

func readMemberGuildSupportItemExpired(ctx *gin.Context) {
	// there is no request body
	session := ctx.MustGet("session").(*userdata.Session)

	user_info_trigger.ReadMemberGuildSupportItemExpired(session)

	common.JsonResponse(ctx, &response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	server.AddHandler("/", "POST", "/infoTrigger/readMemberGuildSupportItemExpired", readMemberGuildSupportItemExpired)
}
