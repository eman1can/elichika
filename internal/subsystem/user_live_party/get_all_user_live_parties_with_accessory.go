package user_live_party

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

func GetAllUserLivePartiesWithAccessory(session *userdata.Session, accessoryId int64) []client.UserLiveParty {
	parties := []client.UserLiveParty{}
	err := session.Db.Table("u_live_party").
		Where("user_id = ? AND (user_accessory_id_1 = ? OR user_accessory_id_2 = ? OR user_accessory_id_3 = ? )",
			session.UserId, accessoryId, accessoryId, accessoryId).Find(&parties)
	utils.CheckErr(err)
	return parties
}
