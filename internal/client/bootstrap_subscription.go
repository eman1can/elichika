package client

import "elichika/internal/generic"

type BootstrapSubscription struct {
	ContinueRewards generic.List[SubscriptionContinueReward] `json:"continue_rewards"`
}
