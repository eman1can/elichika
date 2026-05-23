package communication_member

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_custom_background"
	"elichika/internal/subsystem/user_member"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func setTheme(ctx *gin.Context) {
	req := request.SetThemeRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	member := user_member.GetMember(session, req.MemberMasterId)
	member.SuitMasterId = req.SuitMasterId
	member.CustomBackgroundMasterId = req.CustomBackgroundMasterId
	user_member.UpdateMember(session, member)
	user_custom_background.ReadCustomBackground(session, req.CustomBackgroundMasterId)

	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	server.AddHandler("/", "POST", "/communicationMember/setTheme", setTheme)
}
