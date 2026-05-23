package response

import (
	"elichika/internal/client"
	"elichika/internal/generic"
)

type FetchLoveRankingResponse struct {
	LoveRankingData generic.Array[client.LoveRankingData] `json:"love_ranking_data"`
	MyRankingOrder  generic.Nullable[int32]               `json:"my_ranking_order"`
}
