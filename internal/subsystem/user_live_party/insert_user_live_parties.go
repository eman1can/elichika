package user_live_party

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
)

func InsertUserLiveParties(session *userdata.Session, parties []client.UserLiveParty) {
	for _, party := range parties {
		UpdateUserLiveParty(session, party)
	}
}
