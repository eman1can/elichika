package user_status

import (
	"elichika/internal/subsystem/user_content"
	"elichika/internal/userdata"
)

func AddUserAccessoryLimit(session *userdata.Session, accessoryLimit int32) {
	user_content.OverflowCheckedAdd(&session.UserStatus.AccessoryBoxLimit, &accessoryLimit)
}
