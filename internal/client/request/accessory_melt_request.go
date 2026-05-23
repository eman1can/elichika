package request

import "elichika/internal/generic"

type AccessoryMeltRequest struct {
	UserAccessoryIds generic.Array[int64] `json:"user_accessory_ids"`
}
