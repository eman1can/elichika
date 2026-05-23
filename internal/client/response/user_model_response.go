package response

import "elichika/internal/client"

type UserModelResponse struct {
	UserModel *client.UserModel `json:"user_model"`
}
