package response

import "elichika/internal/client"

type LevelUpCardResponse struct {
	UserModelDiff *client.UserModel `json:"user_model_diff"` // is actually named _UserModelDiff
}
