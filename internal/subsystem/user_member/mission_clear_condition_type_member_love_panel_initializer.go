package user_member

import (
	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/subsystem/user_mission"
	"elichika/internal/userdata"
)

func missionClearConditionTypeMemberLovePanelInitializer(session *userdata.Session, userMission client.UserMission) client.UserMission {
	mission := session.Gamedata.Mission[userMission.MissionMId]
	userMission.MissionCount = 0

	for _, masterPanel := range session.Gamedata.MemberLovePanel {
		// Skip other members panels if the mission is only for a specific member
		if mission.MissionClearConditionParam1 != nil && *mission.MissionClearConditionParam1 == *masterPanel.MemberMasterId {
			continue
		}

		panel := GetMemberLovePanel(session, *masterPanel.MemberMasterId, masterPanel.Id)
		if panel.AllUnlocked() {
			userMission.MissionCount++
		}
	}

	userMission.IsCleared = userMission.MissionCount >= mission.MissionClearConditionCount
	return userMission
}

func init() {
	user_mission.AddMissionInitializer(enum.MissionClearConditionTypeMemberLovePanel, missionClearConditionTypeMemberLovePanelInitializer)
}
