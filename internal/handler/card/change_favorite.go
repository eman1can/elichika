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

func changeFavorite(ctx *gin.Context) {
	req := request.ChangeFavoriteRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	userCard := user_card.GetUserCard(session, req.CardMasterId)
	userCard.IsFavorite = req.IsFavorite
	user_card.UpdateUserCard(session, userCard)

	common.JsonResponse(ctx, &response.ChangeFavoriteResponse{
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	server.AddHandler("/", "POST", "/card/changeFavorite", changeFavorite)
}
