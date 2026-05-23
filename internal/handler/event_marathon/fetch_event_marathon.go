//go:build !dev

package event_marathon

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

// response: FetchEventMarathonResponse
// alternative response: RecoverableExceptionResponse

func fetchEventMarathon(ctx *gin.Context) {
	req := request.FetchEventMarathonRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)
	success, failure := event.FetchEventMarathon(session, req.EventId)
	if success != nil {
		common.JsonResponse(ctx, success)
	} else {
		common.AlternativeJsonResponse(ctx, failure)
	}
}

func init() {
	server.AddHandler("/", "POST", "/eventMarathon/fetchEventMarathon", fetchEventMarathon)
}
