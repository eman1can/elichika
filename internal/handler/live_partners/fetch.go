package live_partners

import (
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_social"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

func fetch(ctx *gin.Context) {
	// there's no request body
	session := ctx.MustGet("session").(*userdata.Session)
	common.JsonResponse(ctx, user_social.GetLivePartners(session))
}

func init() {
	server.AddHandler("/", "POST", "/livePartners/fetch", fetch)
}
