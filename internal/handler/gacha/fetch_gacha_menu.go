package gacha

import (
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_gacha"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

func fetchGachaMenu(ctx *gin.Context) {
	// there is no request body
	session := ctx.MustGet("session").(*userdata.Session)

	common.JsonResponse(ctx, &response.FetchGachaMenuResponse{
		GachaList:     user_gacha.GetGachaList(session),
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	server.AddHandler("/", "POST", "/gacha/fetchGachaMenu", fetchGachaMenu)
}
