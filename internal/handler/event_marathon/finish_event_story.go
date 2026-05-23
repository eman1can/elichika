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

// response: UserModelResponse
func finishEventStory(ctx *gin.Context) {
	req := request.FinishEventStoryRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)
	common.JsonResponse(ctx, event.FinishEventStory(session, req.StoryEventMasterId, req.IsAutoMode))
}

func init() {
	server.AddHandler("/", "POST", "/eventMarathon/finishEventStory", finishEventStory)
}
