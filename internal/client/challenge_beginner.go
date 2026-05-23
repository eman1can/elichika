package client

import "elichika/internal/generic"

type ChallengeBeginner struct {
	ChallengeId int32                       `json:"challenge_id"`
	Cells       generic.List[ChallengeCell] `json:"cells"`
}
