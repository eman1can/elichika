package user_tutorial

import (
	"log"

	"elichika/internal/enum"
	"elichika/internal/userdata"
)

func CorePlayableEnd(session *userdata.Session) {
	if session.UserStatus.TutorialPhase != enum.TutorialPhaseCorePlayable {
		log.Panic("Unexpected tutorial phase")
	}
	session.UserStatus.TutorialPhase = enum.TutorialPhaseTimingAdjuster
}
