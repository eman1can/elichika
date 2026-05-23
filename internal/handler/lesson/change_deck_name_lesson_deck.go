package handler

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_lesson_deck"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

// request: ChangeNameLessonDeckRequest
// response: UserModelResponse
// error response: RecoverableExceptionResponse
func changeDeckNameLessonDeck(ctx *gin.Context) {
	req := request.ChangeNameLessonDeckRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	successResponse, failureResponse := user_lesson_deck.SetLessonDeckName(session, req.DeckId, req.DeckName)
	if successResponse != nil {
		common.JsonResponse(ctx, successResponse)
	} else {
		common.AlternativeJsonResponse(ctx, failureResponse)
	}
}

func init() {
	server.AddHandler("/", "POST", "/lesson/changeDeckNameLessonDeck", changeDeckNameLessonDeck)
}
