package navi

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_voice"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func saveUserNaviVoice(ctx *gin.Context) {
	req := request.SaveUserNaviVoiceRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	for _, naviVoiceMasterId := range req.NaviVoiceMasterIds.Slice {
		user_voice.UpdateUserVoice(session, naviVoiceMasterId, false)
	}

	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	server.AddHandler("/", "POST", "/navi/saveUserNaviVoice", saveUserNaviVoice)
}
