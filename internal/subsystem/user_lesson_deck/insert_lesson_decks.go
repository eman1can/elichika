package user_lesson_deck

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
)

func InsertLessonDecks(session *userdata.Session, decks []client.UserLessonDeck) {
	for _, deck := range decks {
		UpdateLessonDeck(session, deck)
	}
}
