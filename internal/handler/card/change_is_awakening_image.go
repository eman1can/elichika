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

func changeIsAwakeningImage(ctx *gin.Context) {
	req := request.ChangeIsAwakeningImageRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	userCard := user_card.GetUserCard(session, req.CardMasterId)
	userCard.IsAwakeningImage = req.IsAwakeningImage
	user_card.UpdateUserCard(session, userCard)

	common.JsonResponse(ctx, response.ChangeIsAwakeningImageResponse{
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	server.AddHandler("/", "POST", "/card/changeIsAwakeningImage", changeIsAwakeningImage)
}
