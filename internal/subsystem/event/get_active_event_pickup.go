//go:build !dev

package event

import (
	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/generic"
	"elichika/internal/userdata"
)

func GetActiveEventPickup(session *userdata.Session) generic.Nullable[client.BootstrapPickupEventInfo] {
	active := session.Gamedata.EventActive
	if active == nil {
		return generic.Nullable[client.BootstrapPickupEventInfo]{}
	}
	result := generic.NewNullable(client.BootstrapPickupEventInfo{
		EventId:   active.EventId,
		StartAt:   active.StartAt,
		ClosedAt:  active.ExpiredAt,
		EndAt:     active.EndAt,
		EventType: active.EventType,
	})
	if session.Time.Unix() < active.ExpiredAt {
		if active.EventType == enum.EventTypeMarathon {
			event := session.Gamedata.EventMarathon[active.EventId]
			result.Value.BoosterItemId = generic.NewNullable(event.BoosterItemId)
		}
	}
	return result
}
