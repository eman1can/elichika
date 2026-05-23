package response

import "elichika/internal/client"

type FinishLiveResponse struct {
	LiveResult    client.LiveResult `json:"live_result"`
	UserModelDiff *client.UserModel `json:"user_model_diff"`
}
