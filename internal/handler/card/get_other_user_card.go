package card

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_card"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func getOtherUserCard(ctx *gin.Context) {
	req := request.GetOtherUserCardRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	// the name of request and response is not consistent for this one, for some reason
	common.JsonResponse(ctx, response.FetchOtherUserCardResponse{
		OtherUserCard: user_card.GetOtherUserCard(session, req.UserId, req.CardMasterId),
	})
}

func init() {
	server.AddHandler("/", "POST", "/card/getOtherUserCard", getOtherUserCard)
}
