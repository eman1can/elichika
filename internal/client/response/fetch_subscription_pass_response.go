package response

import "elichika/internal/generic"

type FetchSubscriptionPassResponse struct {
	BeforeContinueCount generic.Nullable[int32] `json:"before_continue_count"`
}
