package friend

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

// request: UserSearchUserIdRequest
// success response: UserSearchResponse
// error response: FriendRecoverableExceptionResponse
func searchUserId(ctx *gin.Context) {
	req := request.UserSearchUserIdRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	successResponse, failureResponse := user_social.SearchUserId(session, req.UserId)
	if successResponse != nil {
		common.JsonResponse(ctx, successResponse)
	} else {
		common.AlternativeJsonResponse(ctx, failureResponse)
	}
}

func init() {
	server.AddHandler("/", "POST", "/friend/searchUserId", searchUserId)
}
