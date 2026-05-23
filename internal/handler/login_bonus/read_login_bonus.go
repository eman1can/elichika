package login_bonus

import (
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"

	"github.com/gin-gonic/gin"
)

func readLoginBonus(ctx *gin.Context) {
	// this doesn't need to do anything, at least with this way of handling things
	// reqBody := ctx.Get("reqBody").(json.RawMessage)
	// req := request.ReadLoginBonusRequest{}
	// err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	// utils.CheckErr(err)
	common.JsonResponse(ctx, &response.EmptyResponse{})
}

func init() {
	server.AddHandler("/", "POST", "/loginBonus/readLoginBonus", readLoginBonus)
}
