//go:build !dev

package event_mining

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/event"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

// response: FetchEventMiningResponse
// alternative response: RecoverableExceptionResponse

func fetchEventMining(ctx *gin.Context) {
	req := request.FetchEventMiningRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)
	success, failure := event.FetchEventMining(session, req.EventId)
	if success != nil {
		common.JsonResponse(ctx, success)
	} else {
		common.AlternativeJsonResponse(ctx, failure)
	}
}

func init() {
	server.AddHandler("/", "POST", "/eventMining/fetchEventMining", fetchEventMining)
}
