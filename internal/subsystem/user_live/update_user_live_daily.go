package user_live

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

func UpdateUserLiveDaily(session *userdata.Session, liveDaily client.LiveDaily) {
	affected, err := session.Db.Table("u_live_daily").Where("user_id = ? AND live_daily_master_id = ?",
		session.UserId, liveDaily.LiveDailyMasterId).AllCols().Update(liveDaily)
	utils.CheckErr(err)
	if affected == 0 {
		userdata.GenericDatabaseInsert(session, "u_live_daily", liveDaily)
	}
}
