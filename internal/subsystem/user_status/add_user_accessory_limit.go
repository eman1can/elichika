package user_status

import "elichika/internal/userdata"

func AddUserAccessoryLimit(session *userdata.Session, accessoryLimit int32) {
	session.UserStatus.AccessoryBoxAdditional += accessoryLimit
}
