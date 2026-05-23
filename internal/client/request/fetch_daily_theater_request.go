package request

import "elichika/internal/generic"

type FetchDailyTheaterRequest struct {
	DailyTheaterId generic.Nullable[int32] `json:"daily_theater_id"`
}
