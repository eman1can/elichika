package login

import (
	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/locale"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_account"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"encoding/json"
	// "net/http"

	"github.com/gin-gonic/gin"
)

func startup(ctx *gin.Context) {
	req := request.StartupRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	resp := response.StartupResponse{}
	resp.UserId = int32(user_account.CreateNewAccount(ctx, -1, ""))

	session := userdata.GetSession(ctx, resp.UserId)
	defer session.Close()
	resp.AuthorizationKey = session.EncodedAuthorizationKey(req.Mask)
	// note that this use a different key than the common one
	ctx.Set("sign_key", ctx.MustGet("locale").(*locale.Locale).StartupKey)
	common.JsonResponse(ctx, &resp)
}

func init() {
	server.AddHandler("/", "POST", "/login/startup", startup)
}
