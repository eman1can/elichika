package user

import (
	"encoding/json"
	"net/http"

	"elichika/internal/gamedata"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_story_event_history"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

type WebUIListEventStoryRequest struct {
	Language string `form:"l" json:"l"`
	EventId  int32  `form:"id" json:"id"`
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
	dictionary := gamedata.DictionaryByLanguage(req.Language)

	for _, storyEvent := range gamedata.Instance.StoryEventHistory {
		if req.EventId != storyEvent.EventMasterId {
			continue
		}

		story := user_story_event_history.GetEventStory(session, storyEvent.StoryEventId)
		resp = append(resp, WebUIStoryEventEntry{
			StoryEventId:   storyEvent.StoryEventId,
			Title:          dictionary.Resolve(storyEvent.Title),
			Description:    dictionary.Resolve(storyEvent.Description),
			ImageAssetPath: storyEvent.BannerThumbnailPath,
			StoryNumber:    storyEvent.StoryNumber,
			IsNew:          story.IsNew,
		})
	}

	jsonBytes, err := json.Marshal(resp)
	utils.CheckErr(err)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, string(jsonBytes))
}

func init() {
	server.AddHandler("/webui/user", "GET", "/list_event_story", listEventStory)
}
