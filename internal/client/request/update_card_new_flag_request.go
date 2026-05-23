package request

import "elichika/internal/generic"

type UpdateCardNewFlagRequest struct {
	CardMasterIds generic.Array[int32] `json:"card_master_ids"`
}
