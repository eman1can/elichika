package user_live_deck

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
)

func UpdateUserLiveDeck(session *userdata.Session, liveDeck client.UserLiveDeck) {
	session.UserModel.UserLiveDeckById.Set(liveDeck.UserLiveDeckId, liveDeck)
}
