package story_event_history

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/config"
	"elichika/internal/handler/common"
	"elichika/internal/item"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_content"
	"elichika/internal/subsystem/user_story_event_history"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func unlockStory(ctx *gin.Context) {
	req := request.UnlockStoryEventHistoryRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	user_story_event_history.UnlockEventStory(session, req.EventStoryMasterId)
	if config.Conf.ResourceConfig().ConsumeMiscItems {
		user_content.RemoveContent(session, item.MemoryKey)
	}

	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	server.AddHandler("/", "POST", "/storyEventHistory/unlockStory", unlockStory)
}
