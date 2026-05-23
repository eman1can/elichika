package user_card

import (
	"elichika/internal/userdata"
	"elichika/internal/userdata/database"
	"elichika/internal/utils"
)

func UpdateUserCardPlayCountStat(session *userdata.Session, userCardPlayCountStat database.UserCardPlayCountStat) {
	affected, err := session.Db.Table("u_card_play_count_stat").
		Where("user_id = ? AND card_master_id = ?", session.UserId, userCardPlayCountStat.CardMasterId).AllCols().Update(&userCardPlayCountStat)
	utils.CheckErr(err)
	if affected == 0 {
		userdata.GenericDatabaseInsert(session, "u_card_play_count_stat", userCardPlayCountStat)
	}
}
