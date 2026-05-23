package user_live_deck

import (
	"reflect"

	"elichika/internal/enum"
	"elichika/internal/generic"
	"elichika/internal/subsystem/user_member"
	"elichika/internal/userdata"
)

func ChangeLiveDeckSuit(session *userdata.Session, deckId, cardIndex, suitMasterId, viewStatus int32) {
	if session.UserStatus.TutorialPhase == enum.TutorialPhaseSuitChange {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseGacha
	}

	userLiveDeck := GetUserLiveDeck(session, deckId)
	reflect.ValueOf(&userLiveDeck).Elem().Field(int(1 + cardIndex + 9)).Set(reflect.ValueOf(generic.NewNullable(suitMasterId)))
	UpdateUserLiveDeck(session, userLiveDeck)

	// Rina-chan board toggle
	if session.Gamedata.Suit[suitMasterId].Member.Id == enum.MemberMasterIdRina {
		RinaChan := user_member.GetMember(session, enum.MemberMasterIdRina)
		RinaChan.ViewStatus = viewStatus
		user_member.UpdateMember(session, RinaChan)
	}

}
