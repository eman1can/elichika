package response

import "elichika/internal/client"

type ChangeFavoriteResponse struct {
	UserModelDiff *client.UserModel `json:"user_model_diff"`
}
