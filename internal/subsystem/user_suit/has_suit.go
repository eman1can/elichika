package user_suit

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

func HasSuit(session *userdata.Session, suitMasterId int32) bool {
	_, exists := session.Gamedata.Suit[suitMasterId]
	if !exists {
		return false
	}
	exists, err := session.Db.Table("u_suit").
		Where("user_id = ? AND suit_master_id = ?", session.UserId, suitMasterId).Exist(&client.UserSuit{})
	utils.CheckErr(err)
	return exists
}
