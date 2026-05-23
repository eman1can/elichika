package response

import (
	"elichika/internal/client"
	"elichika/internal/generic"
)

type FetchTradeResponse struct {
	Trades generic.Array[client.Trade] `json:"trades"` // the name is actually _Trades, for some reason
}
