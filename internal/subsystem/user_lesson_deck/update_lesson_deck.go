package user_lesson_deck

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
)

func UpdateLessonDeck(session *userdata.Session, userLessonDeck client.UserLessonDeck) {
	session.UserModel.UserLessonDeckById.Set(userLessonDeck.UserLessonDeckId, userLessonDeck)
}
