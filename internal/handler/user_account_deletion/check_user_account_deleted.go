package user_account_deletion

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_social"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func checkUserAccountDeleted(ctx *gin.Context) {
	req := request.UserAccountDeletionRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	// response with an empty response, or more precisely UserAccountDeletionRecoverableExceptionResponse if the user exist
	session := ctx.MustGet("session").(*userdata.Session)
	if !user_social.UserExist(session, req.UserId) {
		common.AlternativeJsonResponse(ctx, response.UserAccountDeletionRecoverableExceptionResponse{})
	} else {
		common.JsonResponse(ctx, response.EmptyResponse{})
	}
}

func init() {
	server.AddHandler("/", "POST", "/userAccountDeletion/checkUserAccountDeleted", checkUserAccountDeleted)
}
