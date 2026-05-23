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

func executeMultiTrade(ctx *gin.Context) {
	req := request.ExecuteMultiTradeRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	sentToPresentBox := false
	for _, trade := range req.TradeOrders.Slice {
		if user_trade.ExecuteTrade(session, trade.ProductId, trade.TradeCount) {
			sentToPresentBox = true
		}
	}
	sentToPresentBox = sentToPresentBox || (len(session.UnreceivedContent) > 0)

	session.Finalize()
	common.JsonResponse(ctx, response.ExecuteTradeResponse{
		Trades:           user_trade.GetTrades(session, user_trade.GetTradeTypeByProductId(session, req.TradeOrders.Slice[0].ProductId)),
		IsSendPresentBox: sentToPresentBox,
		UserModelDiff:    &session.UserModel,
	})
}

func init() {
	server.AddHandler("/", "POST", "/trade/executeMultiTrade", executeMultiTrade)
}
