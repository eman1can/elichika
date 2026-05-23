package scene_tips

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_scene_tips"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func saveSceneTipsType(ctx *gin.Context) {
	req := request.SaveSceneTipsRequest1{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	user_scene_tips.SaveUserSceneTips(session, req.SceneTipsType)

	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	server.AddHandler("/", "POST", "/sceneTips/saveSceneTipsType", saveSceneTipsType)
}
