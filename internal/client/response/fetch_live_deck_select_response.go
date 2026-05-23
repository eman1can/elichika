package response

import (
	"elichika/internal/client"
	"elichika/internal/generic"
)

type FetchLiveDeckSelectResponse struct {
	LastPlayLiveDifficultyDeck generic.Nullable[client.LastPlayLiveDifficultyDeck] `json:"last_play_live_difficulty_deck"` // pointer
}
