package request

import (
	"elichika/internal/client"
	"elichika/internal/generic"
)

type ExecuteMultiTradeRequest struct {
	TradeOrders generic.Array[client.TradeOrder] `json:"trade_orders"`
}
