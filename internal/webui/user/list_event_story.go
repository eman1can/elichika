package user

import (
	"net/http"

	"elichika/internal/gamedata"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_story_event_history"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

func listEventStory(ctx *gin.Context) {
	var resp []WebUIStoryChapterEntry

	session := ctx.MustGet("session").(*userdata.Session)
	dictionary := ctx.MustGet("dictionary").(*gamedata.Dictionary)

	for eventId, event := range session.Gamedata.Event {
		entry := WebUIStoryChapterEntry{
			Id:             eventId,
			Title:          dictionary.Resolve(event.Title),
			Description:    dictionary.Resolve(event.Description),
			DisplayOrder:   event.ReleaseOrder,
			ImageAssetPath: *event.BannerNoticeLargeAssetPath,
			Chapters:       make([]WebUIStoryCellEntry, 0),
		}

		for _, storyEvent := range session.Gamedata.StoryEventHistory {
			if storyEvent.EventMasterId != eventId {
				continue
			}

			entry.Chapters = append(entry.Chapters, WebUIStoryCellEntry{
				Id:             storyEvent.StoryEventId,
				Chapter:        storyEvent.StoryNumber,
				Title:          dictionary.Resolve(storyEvent.Title),
				Description:    dictionary.Resolve(storyEvent.Description),
				ImageAssetPath: storyEvent.BannerThumbnailPath,
				IsNew:          !user_story_event_history.IsEventStoryFinished(session, storyEvent.StoryEventId),
				Unlocked:       true,
			})
		}

		resp = append(resp, entry)
	}

	ctx.JSON(http.StatusOK, resp)
}

func init() {
	server.AddHandler("/webui/user", "GET", "/list_event_story", listEventStory)
}
