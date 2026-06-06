package user

import (
	"net/http"

	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/generic"
	"elichika/internal/item"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_present"
	"elichika/internal/subsystem/user_story_main"

	"github.com/gin-gonic/gin"
)

func finishMainStory(ctx *gin.Context) {
	session, masterIds := getMasterIdsSession(ctx)

	if session != nil {
		for _, cellId := range masterIds {
			// TODO: Should we move present awarding to the completion handler?
			if user_story_main.InsertUserStoryMainCell(session, cellId) {
				user_present.AddPresent(session, client.PresentItem{
					Content:          item.StarGem.Amount(10),
					PresentRouteType: enum.PresentRouteTypeStoryMain,
					PresentRouteId:   generic.NewNullable(cellId),
				})
			}
		}

		session.Finalize()
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

func init() {
	server.AddHandler("/webui/user", "POST", "/finish_main_story", finishMainStory)
}
