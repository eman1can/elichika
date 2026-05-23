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

func executeTrade(ctx *gin.Context) {
	req := request.ExecuteTradeRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)
	// TODO(trade): This part can be wrong if user_trade.GetTrades require the session
	// this only decide whether there's a text saying that things were sent to present box
	sentToPresentBox := user_trade.ExecuteTrade(session, req.ProductId, req.TradeCount)
	sentToPresentBox = sentToPresentBox || (len(session.UnreceivedContent) > 0)

	session.Finalize()
	common.JsonResponse(ctx, response.ExecuteTradeResponse{
		Trades:           user_trade.GetTrades(session, user_trade.GetTradeTypeByProductId(session, req.ProductId)),
		IsSendPresentBox: sentToPresentBox,
		UserModelDiff:    &session.UserModel,
	})
}

func init() {
	server.AddHandler("/", "POST", "/trade/executeTrade", executeTrade)
}
