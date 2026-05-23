package member_guild

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_member_guild"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func fetchMemberGuildRankingYear(ctx *gin.Context) {
	req := request.FetchMemberGuildRankingYearRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	common.JsonResponse(ctx, response.FetchMemberGuildRankingYearResponse{
		user_member_guild.FetchMemberGuildRankingYear(session, req.Year),
	})
}

func init() {
	server.AddHandler("/", "POST", "/memberGuild/fetchMemberGuildRankingYear", fetchMemberGuildRankingYear)
}
