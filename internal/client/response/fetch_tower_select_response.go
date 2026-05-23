package response

import (
	"elichika/internal/client"
	"elichika/internal/generic"
)

type FetchTowerSelectResponse struct {
	TowerIds      generic.Array[int32] `json:"tower_ids"`
	UserModelDiff *client.UserModel    `json:"user_model_diff"`
}
