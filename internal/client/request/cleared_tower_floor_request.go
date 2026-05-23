package request

import "elichika/internal/generic"

type ClearedTowerFloorRequest struct {
	TowerId    int32                  `json:"tower_id"`
	FloorNo    int32                  `json:"floor_no"`
	IsAutoMode generic.Nullable[bool] `json:"is_auto_mode"`
}
