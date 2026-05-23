package user_accessory

import "elichika/internal/userdata"

func UpdateIsLock(session *userdata.Session, userAccessoryId int64, isLock bool) {
	accessory := GetUserAccessory(session, userAccessoryId)
	accessory.IsLock = isLock
	UpdateUserAccessory(session, accessory)
}
