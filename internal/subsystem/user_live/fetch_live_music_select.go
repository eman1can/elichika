package user_live

import (
	"time"

	"elichika/internal/client"
	"elichika/internal/client/response"
	"elichika/internal/subsystem/event"
	"elichika/internal/subsystem/pickup_info"
	"elichika/internal/userdata"
)

func FetchLiveMusicSelect(session *userdata.Session) response.FetchLiveMusicSelectResponse {
	year, month, day := session.Time.Year(), session.Time.Month(), session.Time.Day()
	tomorrow := time.Date(year, month, day+1, 0, 0, 0, 0, session.Time.Location()).Unix()

	weekday := int32(session.Time.Weekday())
	if weekday == 0 {
		weekday = 7
	}

	resp := response.FetchLiveMusicSelectResponse{
		WeekdayState: client.WeekdayState{
			Weekday:       weekday,
			NextWeekdayAt: tomorrow,
		},
		UserModelDiff: &session.UserModel,
	}

	for _, liveDaily := range session.Gamedata.LiveDaily {
		if liveDaily.Weekday != weekday {
			continue
		}
		resp.LiveDailyList.Append(GetUserLiveDaily(session, liveDaily.Id))
	}

	resp.LiveEventCommonInfo = event.GetLiveEventCommonInfo(session)
	resp.PickupInfo = pickup_info.GetBootstrapPickupInfo(session)
	return resp
}
