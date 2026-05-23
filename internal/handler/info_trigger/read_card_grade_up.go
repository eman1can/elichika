package info_trigger

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_info_trigger"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func readCardGradeUp(ctx *gin.Context) {
	req := request.ReadInfoTriggerRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	user_info_trigger.RemoveTriggerCardGradeUp(session, req.TriggerId)

	common.JsonResponse(ctx, &response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	server.AddHandler("/", "POST", "/infoTrigger/readCardGradeUp", readCardGradeUp)
}
