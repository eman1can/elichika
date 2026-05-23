package response

import "elichika/internal/client"

type UpdateCardNewFlagResponse struct {
	UserModelDiff *client.UserModel `json:"user_model_diff"`
}
