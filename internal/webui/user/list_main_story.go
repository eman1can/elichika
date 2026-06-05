package user

import (
	"net/http"

	"elichika/internal/gamedata"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_story_main"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

type WebUIMainStoryChapterEntry struct {
	Id             int32  `json:"id"`
	Title          string `json:"title"`
	ImageAssetPath string `json:"image_asset_path"`
	IsNew          bool   `json:"is_new"`
}

func listMainStory(ctx *gin.Context) {
	var resp []WebUIMainStoryChapterEntry

	session := ctx.MustGet("session").(*userdata.Session)
	dictionary := ctx.MustGet("dictionary").(*gamedata.Dictionary)

	for _, storyMainChapter := range session.Gamedata.StoryMainChapter {
		resp = append(resp, WebUIMainStoryChapterEntry{
			Id:             storyMainChapter.Id,
			Title:          dictionary.Resolve(storyMainChapter.Title),
			ImageAssetPath: storyMainChapter.ThumbnailAssetPath,
			IsNew:          !user_story_main.IsStoryFinished(session, storyMainChapter.Id),
		})
	}

	ctx.JSON(http.StatusOK, resp)
}

func init() {
	server.AddHandler("/webui/user", "GET", "/list_main_story", listMainStory)
}
