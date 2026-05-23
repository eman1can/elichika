package request

import "elichika/internal/generic"

type CheerMemberGuildRequest struct {
	CheerItemAmount generic.Nullable[int32] `json:"cheer_item_amount"`
}
