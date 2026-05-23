package request

import "elichika/internal/generic"

type SaveLessonDeckRequest struct {
	DeckId        int32                                              `json:"deck_id"`
	CardMasterIds generic.Dictionary[int32, generic.Nullable[int32]] `json:"card_master_ids"`
}
