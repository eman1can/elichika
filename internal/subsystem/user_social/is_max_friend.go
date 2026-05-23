package user_social

import "elichika/internal/userdata"

func IsMaxFriend(session *userdata.Session) bool {
	friendCount := GetFriendCount(session, session.UserId)
	return friendCount >= session.Gamedata.UserRank[session.UserStatus.Rank].FriendLimit
}
