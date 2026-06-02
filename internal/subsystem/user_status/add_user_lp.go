package user_status

import (
	"elichika/internal/enum"
	"elichika/internal/userdata"
)

// AddUserLivePoints
// If user has full or over full live points, then LivePointBroken holds the user's LP amount
// Otherwise, LivePointBroken holds the max amount  of LP the user can hold
func AddUserLivePoints(session *userdata.Session, lp int32) {
	maxLp := session.Gamedata.UserRank[session.UserStatus.Rank].MaxLp
	currentLp := getUserLivePoints(session)

	currentLp = min(10000, currentLp+lp)
	if currentLp >= maxLp {
		session.UserStatus.LivePointBroken = currentLp
		session.UserStatus.LivePointFullAt = session.Time.Unix()
	} else {
		livePointsRecoverAt := session.Gamedata.ConstantInt[enum.ConstantIntLivePointRecoverAt]
		session.UserStatus.LivePointBroken = maxLp
		session.UserStatus.LivePointFullAt = session.Time.Unix() + int64((maxLp-currentLp)*livePointsRecoverAt)
	}
}
