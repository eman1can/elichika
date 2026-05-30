//go:build dev

package event_marathon

import (
	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/config"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/event"
	"elichika/internal/userdata"
	"elichika/internal/utils"
	"elichika/internal/webui/event_marathon_dev"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

// response: FetchEventMarathonResponse
// alternative response: RecoverableExceptionResponse

func fetchEventMarathon(ctx *gin.Context) {
	req := request.FetchEventMarathonRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)
	if config.DeveloperMode == config.DeveloperModeEventMarathonDev {
		// special case for developer mode
		response := &response.FetchEventMarathonResponse{
			EventMarathonTopStatus: event_marathon_dev.TopStatus,
			UserModelDiff:          &session.UserModel,
		}
		common.JsonResponse(ctx, response)
		return
	}
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
