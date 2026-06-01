package user

import (
	"net/http"

	"elichika/internal/server"
	"elichika/internal/subsystem/user_story_main"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

func finishMainStory(ctx *gin.Context) {
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
	server.AddHandler("/webui/user", "POST", "/finish_main_story", finishMainStory)
}
