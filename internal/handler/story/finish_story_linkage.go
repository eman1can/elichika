package story

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_story_linkage"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func finishStoryLinkage(ctx *gin.Context) {
	req := request.AddStoryLinkageRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	if req.IsAutoMode.HasValue {
		session.UserStatus.IsAutoMode = req.IsAutoMode.Value
	}
	user_story_linkage.InsertUserStoryLinkage(session, req.CellId)

	common.JsonResponse(ctx, &response.AddStoryLinkageResponse{
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	server.AddHandler("/", "POST", "/story/finishStoryLinkage", finishStoryLinkage)
}
