package response

import "elichika/internal/generic"

type GetClearedPlatformAchievementResponse struct {
	ClearedIds generic.Array[string] `json:"cleared_ids"`
}
