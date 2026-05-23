package story

import (
	"encoding/json"

	"elichika/internal/client"
	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/enum"
	"elichika/internal/generic"
	"elichika/internal/handler/common"
	"elichika/internal/item"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_present"
	"elichika/internal/subsystem/user_story_main"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func finishUserStoryMain(ctx *gin.Context) {
	req := request.StoryMainRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	if req.IsAutoMode.HasValue {
		session.UserStatus.IsAutoMode = req.IsAutoMode.Value
	}
	resp := response.StoryMainResponse{
		UserModelDiff: &session.UserModel,
	}

	if user_story_main.InsertUserStoryMain(session, req.CellId) { // newly inserted story, award some gem
		resp.FirstClearReward.Append(item.StarGem.Amount(10))
		user_present.AddPresent(session, client.PresentItem{
			Content:          item.StarGem.Amount(10),
			PresentRouteType: enum.PresentRouteTypeStoryMain,
			PresentRouteId:   generic.NewNullable(req.CellId),
		})
	}
	if req.MemberId.HasValue { // has a member -> select member thingy
		user_story_main.UpdateUserStoryMainSelected(session, req.CellId, req.MemberId.Value)
	}

	common.JsonResponse(ctx, &resp)
}

func init() {
	server.AddHandler("/", "POST", "/story/finishUserStoryMain", finishUserStoryMain)
}
