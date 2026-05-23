package user_member

import (
	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/subsystem/user_mission"
	"elichika/internal/userdata"
)

func missionClearConditionTypeCountLoveLevelInitializer(session *userdata.Session, userMission client.UserMission) client.UserMission {
	mission := session.Gamedata.Mission[userMission.MissionMId]

	userMission.MissionCount = 0
	for memberId := range session.Gamedata.Member {
		member := GetMember(session, memberId)
		if member.LoveLevel >= *mission.MissionClearConditionParam1 {
			userMission.MissionCount++
		}
	}
	userMission.IsCleared = userMission.MissionCount >= mission.MissionClearConditionCount
	return userMission
}

func init() {
	user_mission.AddMissionInitializer(enum.MissionClearConditionTypeCountLoveLevel, missionClearConditionTypeCountLoveLevelInitializer)
}
