package client

import "elichika/internal/generic"

type ProfileLiveDifficulty struct {
	LiveDifficultyMasterId generic.Nullable[int32] `json:"live_difficulty_master_id"`
	Score                  int32                   `json:"score"`
}
