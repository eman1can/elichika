package user_live_difficulty

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
)

func GetUserLiveDifficulty(session *userdata.Session, liveDifficultyId int32) client.UserLiveDifficulty {
	return GetOtherUserLiveDifficulty(session, session.UserId, liveDifficultyId)
}
