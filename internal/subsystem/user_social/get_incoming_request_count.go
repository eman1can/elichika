package user_social

import (
	"elichika/internal/enum"
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

func GetIncomingRequestCount(session *userdata.Session, userId int32) int32 {
	count, err := session.Db.Table("u_friend_status").Where("user_id = ? AND request_status = ?", userId,
		enum.FriendRequestStatusNone).Count()
	utils.CheckErr(err)
	return int32(count)
}
