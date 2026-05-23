package response

import (
	"elichika/internal/client"
	"elichika/internal/generic"
)

type FetchMissionResponse struct {
	MissionMasterIdList generic.List[int32] `json:"mission_master_id_list"`
	UserModel           *client.UserModel   `json:"user_model"`
}
