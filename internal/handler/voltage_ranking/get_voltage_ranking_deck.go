package voltage_ranking

import (
	"elichika/internal/client/request"
	"elichika/internal/server"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	// "elichika/internal/client/response"
	// "elichika/internal/enum"
	// "elichika/internal/generic"
	"encoding/json"

	"elichika/internal/handler/common"
	"elichika/internal/subsystem/voltage_ranking"

	"github.com/gin-gonic/gin"
)

func getVoltageRankingDeck(ctx *gin.Context) {
	req := request.GetVoltageRankingDeckRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	common.JsonResponse(ctx, voltage_ranking.GetVoltageRankingDeckResponse(session, req.LiveDifficultyId, req.UserId))
}

func init() {
	server.AddHandler("/", "POST", "voltageRanking/getVoltageRankingDeck", getVoltageRankingDeck)
}
