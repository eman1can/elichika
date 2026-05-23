package event_marathon_ranking

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_event/marathon"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

// response: FetchEventMarathonRankingResponse
// alternative response: RecoverableExceptionResponse
func fetchEventMarathonRanking(ctx *gin.Context) {
	req := request.FetchEventMarathonRankingRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	success, failure := marathon.FetchEventMarathonRanking(session, req.EventId)
	if success != nil {
		common.JsonResponse(ctx, success)
	} else {
		common.AlternativeJsonResponse(ctx, failure)
	}
}

func init() {
	server.AddHandler("/", "POST", "/eventMarathonRanking/fetchEventMarathonRanking", fetchEventMarathonRanking)
}
