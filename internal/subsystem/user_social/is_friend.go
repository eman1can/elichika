package user_social

import (
	"elichika/internal/enum"
	"elichika/internal/userdata"
)

func IsFriend(session *userdata.Session, otherUserId int32) bool {
	return GetUserFriendStatus(session, otherUserId).RequestStatus == enum.FriendRequestStatusFriend
}
