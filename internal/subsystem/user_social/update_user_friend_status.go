package user_social

import (
	"log"

	"elichika/internal/userdata"
	"elichika/internal/userdata/database"
	"elichika/internal/utils"
)

func UpdateUserFriendStatus(session *userdata.Session, status database.UserFriendStatus) {
	// update only, DO NOT not insert
	affected, err := session.Db.Table("u_friend_status").Where("user_id = ? AND other_user_id = ?",
		status.UserId, status.OtherUserId).AllCols().Update(&status)
	utils.CheckErr(err)
	if affected == 0 {
		log.Panic("friend status doesn't exist")
	}
}
