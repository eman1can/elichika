package user_social

import (
	"elichika/internal/client"
	"elichika/internal/generic"
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

func UserExist(session *userdata.Session, userId int32) bool {
	exist, err := session.Db.Table("u_status").Exist(
		&generic.UserIdWrapper[client.UserStatus]{
			UserId: userId,
		})
	utils.CheckErr(err)
	return exist
}
