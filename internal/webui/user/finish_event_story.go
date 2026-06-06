package user

import (
	"net/http"

	"elichika/internal/server"
	"elichika/internal/subsystem/user_story_event_history"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

type WebUIFinishEventStoryRequest struct {
	MasterIds []int32 `form:"master_ids" json:"master_ids"`
}

func finishEventStory(ctx *gin.Context) {
	var req WebUIFinishEventStoryRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := ctx.MustGet("session").(*userdata.Session)

	for _, masterId := range req.MasterIds {
		user_story_event_history.UnlockEventStory(session, masterId)
		// TODO: Add completion rewards to present box
	}

	session.Finalize()
	ctx.JSON(http.StatusOK, gin.H{})
}

func init() {
	server.AddHandler("/webui/user", "POST", "/finish_event_story", finishEventStory)
}
