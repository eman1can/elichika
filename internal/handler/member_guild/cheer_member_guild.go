package member_guild

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_member_guild"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func cheerMemberGuild(ctx *gin.Context) {
	req := request.CheerMemberGuildRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	resp := user_member_guild.CheerMemberGuild(session, req.CheerItemAmount)
	session.Finalize()
	resp.MemberGuildTopStatus = user_member_guild.FetchMemberGuildTop(session).MemberGuildTopStatus
	common.JsonResponse(ctx, &resp)
}

func init() {
	server.AddHandler("/", "POST", "/memberGuild/cheerMemberGuild", cheerMemberGuild)
}
