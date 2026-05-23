package response

import "elichika/internal/client"

type StartLiveMvResponse struct {
	UniqId        int64             `json:"uniq_id"`
	UserModelDiff *client.UserModel `json:"user_model_diff"`
}
