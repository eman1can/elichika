package response

import "elichika/internal/client"

type AddStoryLinkageResponse struct {
	UserModelDiff        *client.UserModel `json:"user_model_diff"`
	HasAdditionalRewards bool              `json:"has_additional_rewards"`
}
