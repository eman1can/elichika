package story

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_story_main"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func saveBrowseStoryMainDigestMovie(ctx *gin.Context) {
	req := request.SaveBrowseStoryMainDigestMovieRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)
	user_story_main.InsertUserStoryMainPartDigestMovie(session, req.PartId)

	common.JsonResponse(ctx, &response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	server.AddHandler("/", "POST", "/story/saveBrowseStoryMainDigestMovie", saveBrowseStoryMainDigestMovie)
}
