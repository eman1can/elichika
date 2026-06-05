package user

import (
	"net/http"

	"elichika/internal/gamedata"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_story_event_history"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

type WebUIListEventStoryRequest struct {
	EventId int32 `form:"id" json:"id"`
}

type WebUIStoryEventEntry struct {
	StoryEventId   int32  `json:"story_event_id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	ImageAssetPath string `json:"image_asset_path"`
	StoryNumber    int32  `json:"chapter"`
	IsNew          bool   `json:"is_new"`
}

func listEventStory(ctx *gin.Context) {
	var req WebUIListEventStoryRequest
	var resp []WebUIStoryEventEntry

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := ctx.MustGet("session").(*userdata.Session)
	dictionary := ctx.MustGet("dictionary").(*gamedata.Dictionary)

	for _, storyEvent := range session.Gamedata.StoryEventHistory {
		if req.EventId != storyEvent.EventMasterId {
			continue
		}

		resp = append(resp, WebUIStoryEventEntry{
			StoryEventId:   storyEvent.StoryEventId,
			Title:          dictionary.Resolve(storyEvent.Title),
			Description:    dictionary.Resolve(storyEvent.Description),
			ImageAssetPath: storyEvent.BannerThumbnailPath,
			StoryNumber:    storyEvent.StoryNumber,
			IsNew:          !user_story_event_history.IsEventStoryFinished(session, storyEvent.StoryEventId),
		})
	}

	ctx.JSON(http.StatusOK, resp)
}

func init() {
	server.AddHandler("/webui/user", "GET", "/list_event_story", listEventStory)
}
