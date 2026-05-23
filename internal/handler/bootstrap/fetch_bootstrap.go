package bootstrap

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_bootstrap"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func fetchBootstrap(ctx *gin.Context) {
	req := request.FetchBootstrapRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	resp := user_bootstrap.FetchBootstrap(session, req)

	common.JsonResponse(ctx, resp)
}

func init() {
	server.AddHandler("/", "POST", "/bootstrap/fetchBootstrap", fetchBootstrap)
}
