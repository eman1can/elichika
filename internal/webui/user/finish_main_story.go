package user

import (
	"net/http"

	"elichika/internal/gamedata"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_story_main"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

type WebUIFinishMainStoryRequest struct {
	MainStoryMasterIds []int32 `json:"main_story_master_ids"`
}

func finishMainStory(ctx *gin.Context) {
	var req WebUIFinishMainStoryRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := ctx.MustGet("session").(*userdata.Session)

	for _, mainStoryMasterId := range req.MainStoryMasterIds {
		if masterMainStory, ok := gamedata.Instance.StoryMainChapter[mainStoryMasterId]; ok {
			if finished := user_story_main.IsStoryFinished(session, mainStoryMasterId); finished {
				continue
			}

			for _, cell := range masterMainStory.Cells {
				user_story_main.InsertUserStoryMain(session, cell)
			}
		}
	}

	session.Finalize()
	ctx.JSON(http.StatusOK, gin.H{})
}

func init() {
	server.AddHandler("/webui/user", "POST", "/finish_main_story", finishMainStory)
}
