package user_live_deck

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
)

func InsertUserLiveDecks(session *userdata.Session, decks []client.UserLiveDeck) {
	for _, deck := range decks {
		UpdateUserLiveDeck(session, deck)
	}
}
