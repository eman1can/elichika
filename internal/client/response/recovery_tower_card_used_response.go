package response

import (
	"elichika/internal/client"
	"elichika/internal/generic"
)

type RecoveryTowerCardUsedResponse struct {
	TowerCardUsedCountRows generic.List[client.TowerCardUsedCount] `json:"tower_card_used_count_rows"`
	UserModelDiff          *client.UserModel                       `json:"user_model_diff"`
}
