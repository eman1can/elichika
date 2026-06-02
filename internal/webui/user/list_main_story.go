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

type WebUIMainStoryChapterEntry struct {
	gamedata.StoryMainChapter
	IsNew bool `xorm:"-" json:"is_new"`
}

type WebUIListMainStoryResponse struct {
	Chapters []WebUIMainStoryChapterEntry `json:"chapters"`
}

func listMainStory(ctx *gin.Context) {
	var resp WebUIListMainStoryResponse

	session := ctx.MustGet("session").(*userdata.Session)

	for _, masterChapter := range gamedata.Instance.StoryMainChapter {
		chapter := user_story_main.IsStoryFinished(session, masterChapter.Id)
		resp.Chapters = append(resp.Chapters, WebUIMainStoryChapterEntry{
			StoryMainChapter: *masterChapter,
			IsNew:            chapter,
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
