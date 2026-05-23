package response

import "elichika/internal/client"

type ChangeIsAwakeningImageResponse struct {
	UserModelDiff *client.UserModel `json:"user_model_diff"`
}
