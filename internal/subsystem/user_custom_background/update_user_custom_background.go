package user_custom_background

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
)

func UpdateUserCustomBackground(session *userdata.Session, userCustomBackground client.UserCustomBackground) {
	session.UserModel.UserCustomBackgroundById.Set(userCustomBackground.CustomBackgroundMasterId, userCustomBackground)
}
