package voltage_ranking

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/voltage_ranking"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func getVoltageRanking(ctx *gin.Context) {
	req := request.GetVoltageRankingRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	common.JsonResponse(ctx, voltage_ranking.GetVoltageRankingResponse(session, req.LiveDifficultyId))
}

func init() {
	server.AddHandler("/", "POST", "voltageRanking/getVoltageRanking", getVoltageRanking)
}
