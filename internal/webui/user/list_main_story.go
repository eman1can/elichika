package user

import (
	"net/http"

	"elichika/internal/gamedata"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_story_main"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

func listMainStory(ctx *gin.Context) {
	var resp []WebUIStoryChapterEntry

	session := ctx.MustGet("session").(*userdata.Session)
	dictionary := ctx.MustGet("dictionary").(*gamedata.Dictionary)

	for _, storyMainChapter := range session.Gamedata.StoryMainChapter {
		entry := WebUIStoryChapterEntry{
			Id:             storyMainChapter.Id,
			Title:          dictionary.Resolve(storyMainChapter.Title),
			Description:    dictionary.Resolve(storyMainChapter.Description),
			DisplayOrder:   storyMainChapter.Id,
			ImageAssetPath: storyMainChapter.ThumbnailAssetPath,
			Chapters:       make([]WebUIStoryCellEntry, 0),
		}

		for _, cellId := range storyMainChapter.Cells {
			cell := session.Gamedata.StoryMainChapterCell[cellId]
			chapter := WebUIStoryCellEntry{
				Id:          cellId,
				Title:       dictionary.Resolve(cell.Title),
				Description: dictionary.Resolve(cell.Description),
				Chapter:     cell.DisplayOrder,
				IsNew:       !user_story_main.HasStoryMainCell(session, cellId),
				Unlocked:    true,
			}

			if cell.ThumbnailAssetPath != nil {
				chapter.ImageAssetPath = *cell.ThumbnailAssetPath
			} else {
				live := session.Gamedata.LiveDifficulty[*cell.LiveDifficultyId].Live
				chapter.ImageAssetPath = live.JacketAssetPath
				chapter.Title = dictionary.Resolve(live.Name)
			}

			entry.Chapters = append(entry.Chapters, chapter)
		}

		resp = append(resp, entry)
	}

	ctx.JSON(http.StatusOK, resp)
}

func init() {
	server.AddHandler("/webui/user", "GET", "/list_main_story", listMainStory)
}
