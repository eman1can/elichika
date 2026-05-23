package response

import "elichika/internal/client"

type FetchEventMiningResponse struct {
	EventMiningTopStatus client.EventMiningTopStatus `json:"event_mining_top_status"`
	UserModelDiff        *client.UserModel           `json:"user_model_diff"`
}
