package request

import "elichika/internal/generic"

type FetchLoveRankingRequest struct {
	LoveRankingType int32                   `json:"love_ranking_type" enum:"LoveRankingType"`
	Condition       int32                   `json:"condition" enum:"LoveRankingConditionType"`
	RankingOrder    generic.Nullable[int32] `json:"ranking_order"`
}
