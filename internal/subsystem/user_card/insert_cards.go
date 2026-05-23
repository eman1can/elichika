package user_card

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
)

// insert all the cards
func InsertCards(session *userdata.Session, cards []client.UserCard) {
	for _, card := range cards {
		session.UserModel.UserCardByCardId.Set(card.CardMasterId, card)
	}
}
