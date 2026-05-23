package live_deck

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_live_deck"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

// request: ChangeNameLiveDeckRequest
// response: UserModelResponse
// error response: RecoverableExceptionResponse
func changeDeckNameLiveDeck(ctx *gin.Context) {
	req := request.ChangeNameLiveDeckRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	successResponse, failureResponse := user_live_deck.SetLiveDeckName(session, req.DeckId, req.DeckName)
	if successResponse != nil {
		common.JsonResponse(ctx, successResponse)
	} else {
		common.AlternativeJsonResponse(ctx, failureResponse)
	}
}

func init() {
	server.AddHandler("/", "POST", "/liveDeck/changeDeckNameLiveDeck", changeDeckNameLiveDeck)
}
