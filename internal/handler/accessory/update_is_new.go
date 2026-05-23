package accessory

import (
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_accessory"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

func updateIsNew(ctx *gin.Context) {
	user_accessory.ClearIsNewFlags(ctx.MustGet("session").(*userdata.Session))

	common.JsonResponse(ctx, &response.EmptyResponse{})
}

func init() {
	server.AddHandler("/", "POST", "/accessory/updateIsNew", updateIsNew)
}
