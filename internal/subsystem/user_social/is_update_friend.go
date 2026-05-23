package user_social

import (
	"elichika/internal/userdata"
	"elichika/internal/userdata/database"
	"elichika/internal/utils"
)

func IsUpdateFriend(session *userdata.Session) bool {
	exist, err := session.Db.Table("u_friend_status").Where("user_id = ? AND is_new != 0", session.UserId).Exist(&database.UserFriendStatus{})
	utils.CheckErr(err)
	return exist
}
