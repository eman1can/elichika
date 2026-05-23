package request

import "elichika/internal/generic"

type CancelFriendApplyRequest struct {
	UserIds generic.Array[int32] `json:"user_ids"`
	IsMass  bool                 `json:"is_mass"`
}
