package event

import (
	"log"

	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/generic"
	"elichika/internal/userdata"
)

func GetLiveEventCommonInfo(session *userdata.Session) generic.Nullable[client.LiveEventCommonInfo] {
	active := session.Gamedata.EventActive
	if (active == nil) || (active.ExpiredAt <= session.Time.Unix()) {
		return generic.Nullable[client.LiveEventCommonInfo]{}
	}
	result := client.LiveEventCommonInfo{
		EventId:   active.EventId,
		EventType: active.EventType,
		ClosedAt:  active.ExpiredAt,
	}
	if active.EventType == enum.EventTypeMarathon {
		event := session.Gamedata.EventMarathon[active.EventId]
		result.PointBoostContentId = generic.NewNullable(event.BoosterItemId)
	} else if active.EventType == enum.EventTypeMining {
		event := session.Gamedata.EventMining[active.EventId]
		result.EventMusics = event.EventMusics
		// TODO(channel): we might need to insert the channel appeal tex for the song here
		for i := range result.EventMusics.Slice {
			result.EventMusics.Slice[i].EndAt = active.ExpiredAt
		}
	} else {
		log.Panic("not supported")
	}
	return generic.NewNullable(result)
}
