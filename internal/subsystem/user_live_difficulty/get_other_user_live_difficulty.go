package user_live_difficulty

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

func GetOtherUserLiveDifficulty(session *userdata.Session, otherUserId int32, liveDifficultyId int32) client.UserLiveDifficulty {
	userLiveDifficulty := client.UserLiveDifficulty{}
	exist, err := session.Db.Table("u_live_difficulty").
		Where("user_id = ? AND live_difficulty_id = ?", otherUserId, liveDifficultyId).
		Get(&userLiveDifficulty)
	utils.CheckErr(err)
	if !exist {
		userLiveDifficulty.LiveDifficultyId = liveDifficultyId
		userLiveDifficulty.EnableAutoplay = true
		userLiveDifficulty.IsNew = true
	}
	return userLiveDifficulty
}
