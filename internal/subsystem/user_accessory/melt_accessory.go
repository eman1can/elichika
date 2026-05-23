package user_accessory

import (
	"elichika/internal/enum"
	"elichika/internal/subsystem/user_content"
	"elichika/internal/subsystem/user_mission"
	"elichika/internal/userdata"
)

func MeltAccessory(session *userdata.Session, userAccessoryId int64) {
	accessory := GetUserAccessory(session, userAccessoryId)
	user_content.AddContent(session, session.Gamedata.Accessory[accessory.AccessoryMasterId].MeltGroup[accessory.Grade].Reward)
	DeleteUserAccessory(session, userAccessoryId)
	// mission
	user_mission.UpdateProgress(session, enum.MissionClearConditionTypeCountAccessoryMelt, nil, nil,
		user_mission.AddProgressHandler, int32(1))
}
