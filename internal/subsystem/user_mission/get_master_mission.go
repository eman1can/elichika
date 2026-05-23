package user_mission

import (
	"log"

	"elichika/internal/client"
	"elichika/internal/gamedata"
	"elichika/internal/userdata"
)

func GetMasterMission(session *userdata.Session, mission any) *gamedata.Mission {
	switch mission := mission.(type) {
	case client.UserMission:
		return session.Gamedata.Mission[mission.MissionMId]
	case client.UserDailyMission:
		return session.Gamedata.Mission[mission.MissionMId]
	case client.UserWeeklyMission:
		return session.Gamedata.Mission[mission.MissionMId]
	default:
		log.Panic("not supported")
		return nil
	}
}
