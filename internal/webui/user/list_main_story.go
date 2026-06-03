package user

import (
	"encoding/json"
	"net/http"

	"elichika/internal/gamedata"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_story_main"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

type WebUIListMainStoryRequest struct {
	Language string `form:"l" json:"l"`
}

type WebUIMainStoryChapterEntry struct {
	Id             int32  `json:"id"`
	Title          string `json:"title"`
	ImageAssetPath string `json:"image_asset_path"`
	IsNew          bool   `json:"is_new"`
}

func listMainStory(ctx *gin.Context) {
	var req WebUIListMainStoryRequest
	var resp []WebUIMainStoryChapterEntry

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := ctx.MustGet("session").(*userdata.Session)
	dictionary := gamedata.DictionaryByLanguage(req.Language)

	for _, storyMainChapter := range gamedata.Instance.StoryMainChapter {
		finished := user_story_main.IsStoryFinished(session, storyMainChapter.Id)
		resp = append(resp, WebUIMainStoryChapterEntry{
			Id:             storyMainChapter.Id,
			Title:          dictionary.Resolve(storyMainChapter.Title),
			ImageAssetPath: storyMainChapter.ThumbnailAssetPath,
			IsNew:          finished,
		})
	}

	jsonBytes, err := json.Marshal(resp)
	utils.CheckErr(err)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, string(jsonBytes))
}

func init() {
	server.AddHandler("/webui/user", "GET", "/list_main_story", listMainStory)
}
