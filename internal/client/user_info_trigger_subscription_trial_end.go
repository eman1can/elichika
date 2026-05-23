package client

import "elichika/internal/generic"

type UserInfoTriggerSubscriptionTrialEndRow struct {
	TriggerId int64                   `json:"trigger_id"`
	StartAt   generic.Nullable[int64] `json:"start_at"`
}
