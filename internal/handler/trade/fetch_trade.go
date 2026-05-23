package trade

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_trade"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func fetchTrade(ctx *gin.Context) {
	req := request.FetchTradeRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	common.JsonResponse(ctx, response.FetchTradeResponse{
		Trades: user_trade.GetTrades(session, req.TradeType),
	})
}

func init() {
	server.AddHandler("/", "POST", "/trade/fetchTrade", fetchTrade)
}
