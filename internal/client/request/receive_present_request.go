package request

import "elichika/internal/generic"

type ReceivePresentRequest struct {
	Ids generic.List[int32] `json:"ids"`
}
