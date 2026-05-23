package user_live_party

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
)

func UpdateUserLiveParty(session *userdata.Session, liveParty client.UserLiveParty) {
	session.UserModel.UserLivePartyById.Set(liveParty.PartyId, liveParty)
}
