package user_tutorial

import (
	"log"

	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/generic"
	"elichika/internal/subsystem/user_live_deck"
	"elichika/internal/subsystem/user_member"
	"elichika/internal/subsystem/user_training_tree"
	"elichika/internal/userdata"
)

func TutorialSkip(session *userdata.Session, cardWithSuitDict generic.Dictionary[int32, generic.Nullable[int32]], squadDict generic.Dictionary[int32, client.LiveSquad]) {
	switch session.UserStatus.TutorialPhase {
	case enum.TutorialPhaseLovePointUp:
		user_member.TapLovePoint(session, session.UserStatus.FavoriteMemberId)
		fallthrough
	case enum.TutorialPhaseTrainingLevelUp:
		user_training_tree.LevelUpCard(session, session.UserStatus.RecommendCardMasterId, 1)
		fallthrough
	case enum.TutorialPhaseTrainingActivateCell:
		cells := []int32{}
		for id := int32(1); id <= 17; id++ {
			cells = append(cells, id)
		}
		user_training_tree.ActivateTrainingTreeCells(session, session.UserStatus.RecommendCardMasterId, cells)
		fallthrough
	case enum.TutorialPhaseDeckEdit:
		user_live_deck.SaveUserLiveDeck(session, 1, cardWithSuitDict, squadDict)
		fallthrough
	case enum.TutorialPhaseSuitChange:
		user_live_deck.ChangeLiveDeckSuit(session, 1, 1, session.UserStatus.RecommendCardMasterId, enum.MemberViewStatusDefault)
		fallthrough
	case enum.TutorialPhaseGacha:
		session.UserStatus.TutorialPhase = enum.TutorialPhaseFinal
		fallthrough
	case enum.TutorialPhaseFinal:
		PhaseEnd(session)
	default:
		log.Panic("unexpected skip")
	}
}
