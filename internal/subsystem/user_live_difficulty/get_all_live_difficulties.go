package user_live_difficulty

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

func GetAllLiveDifficulties(session *userdata.Session) []client.UserLiveDifficulty {
	records := []client.UserLiveDifficulty{}
	err := session.Db.Table("u_live_difficulty").Where("user_id = ?", session.UserId).
		Find(&records)
	utils.CheckErr(err)
	return records
}
