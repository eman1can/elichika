package user_status

import (
	"elichika/internal/enum"
	"elichika/internal/userdata"
)

func getUserLivePoints(session *userdata.Session) int32 {
	maxLp := session.Gamedata.UserRank[session.UserStatus.Rank].MaxLp
	if session.UserStatus.LivePointBroken >= maxLp {
		return session.UserStatus.LivePointBroken
	}

	livePointsRecoverAt := session.Gamedata.ConstantInt[enum.ConstantIntLivePointRecoverAt]
	timeLeft := int32(session.UserStatus.LivePointFullAt - session.Time.Unix())
	toRecover := timeLeft / livePointsRecoverAt
	if timeLeft%livePointsRecoverAt != 0 {
		toRecover++
	}
	return session.UserStatus.LivePointBroken - toRecover
}

func GetUserLivePoints(session *userdata.Session) int32 {
	return getUserLivePoints(session)
}
