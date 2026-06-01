package user

import (
	"net/http"

	"elichika/internal/server"
	"elichika/internal/subsystem/user_story_main"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

type WebUIFinishEventStoryRequest struct {
	EventMasterIds []int32 `form:"event_master_ids" json:"event_master_ids"`
}

func finishEventStory(ctx *gin.Context) {
	req := WebUIFinishEventStoryRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := ctx.MustGet("session").(*userdata.Session)

	for _, story := range session.Gamedata.StoryMainChapter {
		for _, cell := range story.Cells {
			user_story_main.InsertUserStoryMain(session, cell)
		}
	}

	session.Finalize()
	ctx.JSON(http.StatusOK, gin.H{})
}

func init() {
	server.AddHandler("/webui/user", "POST", "/finish_event_story", finishEventStory)
}
