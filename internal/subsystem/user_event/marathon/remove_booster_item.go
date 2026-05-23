package marathon

import (
	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/subsystem/user_content"
	"elichika/internal/userdata"
)

func RemoveBoosterItem(session *userdata.Session, count int32) {
	event := session.Gamedata.EventActive.GetEventMarathon()
	user_content.RemoveContent(session, client.Content{
		ContentType:   enum.ContentTypeEventMarathonBooster,
		ContentId:     event.BoosterItemId,
		ContentAmount: count,
	})
}
