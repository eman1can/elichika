package tutorial

import (
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_tutorial"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

func phaseEnd(ctx *gin.Context) {
	// there's no request body
	session := ctx.MustGet("session").(*userdata.Session)

	user_tutorial.PhaseEnd(session)

	common.JsonResponse(ctx, &response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	server.AddHandler("/", "POST", "/tutorial/phaseEnd", phaseEnd)
}
