package user_profile

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_social"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func fetchProfile(ctx *gin.Context) {
	req := request.UserProfileRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	common.JsonResponse(ctx, user_social.GetOtherUserProfileResponse(session, req.UserId))
}

func init() {
	server.AddHandler("/", "POST", "/userProfile/fetchProfile", fetchProfile)
}
