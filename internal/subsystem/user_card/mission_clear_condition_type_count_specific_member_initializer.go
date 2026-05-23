package user_card

import (
	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/subsystem/user_mission"
	"elichika/internal/userdata"
)

func missionClearConditionTypeCountSpecificMemberInitializer(session *userdata.Session, userMission client.UserMission) client.UserMission {
	mission := session.Gamedata.Mission[userMission.MissionMId]
	memberId := *mission.MissionClearConditionParam1
	// TODO(mission): This might be a bit slow, but maybe it's fine
	userMission.MissionCount = 0
	for _, card := range session.Gamedata.CardByMemberId[memberId] {
		if userdata.GenericDatabaseExist(session, "u_card", client.UserCard{CardMasterId: card.Id}) {
			userMission.MissionCount++
		}
	}
	userMission.IsCleared = userMission.MissionCount >= mission.MissionClearConditionCount
	return userMission
}

func init() {
	user_mission.AddMissionInitializer(enum.MissionClearConditionTypeCountSpecificMember, missionClearConditionTypeCountSpecificMemberInitializer)
}
