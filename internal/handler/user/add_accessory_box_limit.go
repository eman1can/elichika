package user

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/item"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_content"
	"elichika/internal/subsystem/user_status"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func addAccessoryBoxLimit(ctx *gin.Context) {
	req := request.AddAccessoryBoxLimitRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	req.Count *= 10 // count is the amount of time it is performed, not the amount of slot / gem used

	user_status.AddUserAccessoryLimit(session, req.Count)
	user_content.RemoveContent(session, item.StarGem.Amount(req.Count))

	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	server.AddHandler("/", "POST", "/user/addAccessoryBoxLimit", addAccessoryBoxLimit)
}
