package mining

import (
	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/generic"
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

func FetchUserInfoTriggerEventMiningShowResultRows(session *userdata.Session, result *generic.List[client.UserInfoTriggerEventMiningShowResultRow]) {
	// this show the popup result sequence, only possible when the event is still available
	// the reward is already delivered when the event end to present box, and can't be missed
	active := session.Gamedata.EventActive
	if (active == nil) || (active.EventType != enum.EventTypeMining) {
		return
	}

	var resultTriggers []client.UserInfoTriggerBasic
	err := session.Db.Table("u_info_trigger_basic").Where("user_id = ? AND info_trigger_type = ? AND param_int = ?",
		session.UserId, enum.InfoTriggerTypeEventMiningShowResult, active.EventId).Find(&resultTriggers)
	utils.CheckErr(err)
	for _, trigger := range resultTriggers {
		result.Append(client.UserInfoTriggerEventMiningShowResultRow{
			TriggerId:     trigger.TriggerId,
			EventMiningId: active.EventId,
			ResultAt:      active.ResultAt,
			EndAt:         active.EndAt,
		})
	}
}
