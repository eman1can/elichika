package response

import "elichika/internal/client"

type StartLiveResponse struct {
	Live          client.Live       `json:"live"`
	UserModelDiff *client.UserModel `json:"user_model_diff"`
}
