package take_over

import (
	"encoding/json"
	"fmt"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_authentication"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func updatePassWord(ctx *gin.Context) {
	req := request.UpdatePassWordRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	user_authentication.SetPassWord(session, req.PassWord)

	common.JsonResponse(ctx, &response.UpdatePassWordResponse{
		TakeOverId: fmt.Sprint(session.UserId),
	})
}

func init() {
	server.AddHandler("/", "POST", "/takeOver/updatePassWord", updatePassWord)
}
