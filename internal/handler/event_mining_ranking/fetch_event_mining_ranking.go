package event_mining_ranking

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_event/mining"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

// response: FetchEventMiningRankingResponse
// alternative response: RecoverableExceptionResponse
func fetchEventMiningRanking(ctx *gin.Context) {
	req := request.FetchEventMiningRankingRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	success, failure := mining.FetchEventMiningRanking(session, req.EventId)
	if success != nil {
		common.JsonResponse(ctx, success)
	} else {
		common.AlternativeJsonResponse(ctx, failure)
	}
}

func init() {
	server.AddHandler("/", "POST", "/eventMiningRanking/fetchEventMiningRanking", fetchEventMiningRanking)
}
