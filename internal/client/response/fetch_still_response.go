package response

import "elichika/internal/generic"

type FetchStillResponse struct {
	NewStillList generic.List[int32] `json:"new_still_list"`
}
