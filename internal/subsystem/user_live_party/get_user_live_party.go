package user_live_party

import (
	"log"

	"elichika/internal/client"
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

func GetUserLiveParty(session *userdata.Session, partyId int) client.UserLiveParty {
	liveParty := client.UserLiveParty{}
	exist, err := session.Db.Table("u_live_party").
		Where("user_id = ? AND party_id = ?", session.UserId, partyId).
		Get(&liveParty)
	utils.CheckErr(err)
	if !exist {
		log.Panic("Party doesn't exist")
	}
	return liveParty
}
