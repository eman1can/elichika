package marathon

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

func GetUserEventMarathon(session *userdata.Session) client.UserEventMarathon {
	active := session.Gamedata.EventActive
	ptr, exist := session.UserModel.UserEventMarathonByEventMasterId.Get(active.EventId)
	if exist {
		return *ptr
	}
	userEventMarathon := client.UserEventMarathon{}
	exist, err := session.Db.Table("u_event_marathon").Where("user_id = ? AND event_master_id = ?", session.UserId, active.EventId).
		Get(&userEventMarathon)
	utils.CheckErr(err)
	if !exist {
		userEventMarathon = client.UserEventMarathon{
			EventMasterId:     active.EventId,
			EventPoint:        0,
			OpenedStoryNumber: 1,
			ReadStoryNumber:   0,
		}
	}
	session.UserModel.UserEventMarathonByEventMasterId.Set(active.EventId, userEventMarathon)
	return userEventMarathon
}
