package user_live_party

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

func GetUserLivePartiesWithDeckId(session *userdata.Session, deckId int32) []client.UserLiveParty {
	liveParties := []client.UserLiveParty{}
	err := session.Db.Table("u_live_party").
		Where("user_id = ? AND user_live_deck_id = ?", session.UserId, deckId).
		OrderBy("party_id").Find(&liveParties)
	utils.CheckErr(err)
	return liveParties
}
