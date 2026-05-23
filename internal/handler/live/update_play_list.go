package live

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_play_list"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func updatePlayList(ctx *gin.Context) {
	req := request.UpdatePlayListRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	user_play_list.UpdateUserPlayList(session, req.GroupNum, req.LiveMasterId, req.IsSet)

	common.JsonResponse(ctx, &response.UpdatePlayListResponse{
		IsSuccess:     true,
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	server.AddHandler("/", "POST", "/live/updatePlayList", updatePlayList)
}
