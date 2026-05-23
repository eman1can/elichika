package user_member

import (
	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/subsystem/user_mission"
	"elichika/internal/userdata"
)

func missionClearConditionTypeMemberLovePanelInitializer(session *userdata.Session, userMission client.UserMission) client.UserMission {
	mission := session.Gamedata.Mission[userMission.MissionMId]
	if mission.MissionClearConditionParam1 == nil {
		userMission.MissionCount = 0
		for memberId := range session.Gamedata.Member {
			lovePanel := GetMemberLovePanel(session, memberId)
			cellCount := lovePanel.MemberLovePanelCellIds.Size()
			userMission.MissionCount += int32(cellCount / 5)
		}
	} else {
		lovePanel := GetMemberLovePanel(session, *mission.MissionClearConditionParam1)
		cellCount := lovePanel.MemberLovePanelCellIds.Size()
		userMission.MissionCount = int32(cellCount / 5)
	}
	userMission.IsCleared = userMission.MissionCount >= mission.MissionClearConditionCount
	return userMission
}

func init() {
	user_mission.AddMissionInitializer(enum.MissionClearConditionTypeMemberLovePanel, missionClearConditionTypeMemberLovePanelInitializer)
}
