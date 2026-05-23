package love_ranking

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_love_ranking"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func fetch(ctx *gin.Context) {
	req := request.FetchLoveRankingRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	common.JsonResponse(ctx, user_love_ranking.FetchLoveRanking(session, req.LoveRankingType, req.Condition, req.RankingOrder))
}

func init() {
	server.AddHandler("/", "POST", "/loveRanking/fetch", fetch)
}
