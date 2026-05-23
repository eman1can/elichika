package response

import (
	"elichika/internal/client"
	"elichika/internal/generic"
)

type FetchChallengeBeginnerResponse struct {
	ChallengeBeginner client.ChallengeBeginner `json:"challenge_beginner"`
	CompletedIds      generic.List[int32]      `json:"completed_ids"`
}
