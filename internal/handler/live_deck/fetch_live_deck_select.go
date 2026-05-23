package live_deck

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_live_difficulty"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func fetchLiveDeckSelect(ctx *gin.Context) {
	// return last deck for this song
	req := request.FetchLiveDeckSelectRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	common.JsonResponse(ctx, response.FetchLiveDeckSelectResponse{
		LastPlayLiveDifficultyDeck: user_live_difficulty.GetLastPlayLiveDifficultyDeck(session, req.LiveDifficultyId),
	})
}

func init() {
	server.AddHandler("/", "POST", "/liveDeck/fetchLiveDeckSelect", fetchLiveDeckSelect)
}
