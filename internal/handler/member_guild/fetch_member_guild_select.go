package member_guild

import (
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_member_guild"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

func fetchMemberGuildSelect(ctx *gin.Context) {
	// There is no request body
	session := ctx.MustGet("session").(*userdata.Session)

	common.JsonResponse(ctx, user_member_guild.FetchMemberGuildSelect(session))
}

func init() {
	server.AddHandler("/", "POST", "/memberGuild/fetchMemberGuildSelect", fetchMemberGuildSelect)
}
