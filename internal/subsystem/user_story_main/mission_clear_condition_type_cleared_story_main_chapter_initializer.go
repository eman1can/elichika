package user_story_main

import (
	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/subsystem/user_mission"
	"elichika/internal/userdata"
)

func missionClearConditionTypeClearedStoryMainChapterInitializer(session *userdata.Session, userMission client.UserMission) client.UserMission {
	mission := session.Gamedata.Mission[userMission.MissionMId]
	chapterId := mission.MissionClearConditionCount
	requiredCell := session.Gamedata.StoryMainChapter[chapterId].LastCellId
	if HasStoryMainCell(session, requiredCell) {
		userMission.MissionCount = chapterId
		userMission.IsCleared = true
	} else {
		userMission.MissionCount = 0
		userMission.IsCleared = false
	}
	return userMission
}

func init() {
	user_mission.AddMissionInitializer(enum.MissionClearConditionTypeClearedStoryMainChapter, missionClearConditionTypeClearedStoryMainChapterInitializer)
}
