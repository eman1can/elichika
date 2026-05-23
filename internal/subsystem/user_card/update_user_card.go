package user_card

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
)

func UpdateUserCard(session *userdata.Session, card client.UserCard) {
	session.UserModel.UserCardByCardId.Set(card.CardMasterId, card)
}
