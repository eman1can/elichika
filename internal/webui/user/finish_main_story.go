package user

import (
	"net/http"

	"elichika/internal/server"
	"elichika/internal/subsystem/user_story_main"

	"github.com/gin-gonic/gin"
)

func finishMainStory(ctx *gin.Context) {
	session, masterIds := getMasterIdsSession(ctx)

	if session != nil {
		for _, cell := range masterIds {
			user_story_main.InsertUserStoryMainCell(session, cell)
		}

		// TODO: Add completion rewards to present box

		session.Finalize()
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

func init() {
	server.AddHandler("/webui/user", "POST", "/finish_main_story", finishMainStory)
}
