package mining

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

func GetUserEventMining(session *userdata.Session) client.UserEventMining {
	active := session.Gamedata.EventActive
	ptr, exist := session.UserModel.UserEventMiningByEventMasterId.Get(active.EventId)
	if exist {
		return *ptr
	}
	userEventMining := client.UserEventMining{}
	exist, err := session.Db.Table("u_event_mining").Where("user_id = ? AND event_master_id = ?", session.UserId, active.EventId).
		Get(&userEventMining)
	utils.CheckErr(err)
	if !exist {
		userEventMining = client.UserEventMining{
			EventMasterId:     active.EventId,
			EventPoint:        0,
			EventVoltagePoint: 0,
			OpenedStoryNumber: 1,
			ReadStoryNumber:   0,
		}
	}
	session.UserModel.UserEventMiningByEventMasterId.Set(active.EventId, userEventMining)
	return userEventMining
}
