package member_guild

import (
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_member_guild"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

func fetchMemberGuildRanking(ctx *gin.Context) {
	session := ctx.MustGet("session").(*userdata.Session)
	common.JsonResponse(ctx, user_member_guild.FetchMemberGuildRanking(session))
}

func init() {
	server.AddHandler("/", "POST", "/memberGuild/fetchMemberGuildRanking", fetchMemberGuildRanking)
}
