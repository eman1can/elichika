package request

import "elichika/internal/generic"

type GetPackUrlRequest struct {
	PackNames generic.List[string] `json:"pack_names"`
}
