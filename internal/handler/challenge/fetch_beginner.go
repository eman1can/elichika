package challenge

import (
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_beginner_challenge"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

func fetchBeginner(ctx *gin.Context) {
	// there is no request body
	session := ctx.MustGet("session").(*userdata.Session)

	common.JsonResponse(ctx, user_beginner_challenge.FetchChallengeBeginner(session))
}

func init() {
	server.AddHandler("/", "POST", "/challenge/fetchBeginner", fetchBeginner)
}
