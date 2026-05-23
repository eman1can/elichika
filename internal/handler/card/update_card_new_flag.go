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

func updateCardNewFlag(ctx *gin.Context) {
	req := request.UpdateCardNewFlagRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	for _, cardMasterId := range req.CardMasterIds.Slice {
		card := user_card.GetUserCard(session, cardMasterId)
		card.IsNew = false
		user_card.UpdateUserCard(session, card)
	}

	common.JsonResponse(ctx, response.UpdateCardNewFlagResponse{
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	server.AddHandler("/", "POST", "/card/updateCardNewFlag", updateCardNewFlag)
}
