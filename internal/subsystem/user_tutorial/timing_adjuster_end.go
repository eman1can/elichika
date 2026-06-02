package user_tutorial

import (
	"log"

	"elichika/internal/enum"
	"elichika/internal/userdata"
)

func TimingAdjusterEnd(session *userdata.Session) {
	if session.UserStatus.TutorialPhase != enum.TutorialPhaseTimingAdjuster {
		log.Panic("Unexpected tutorial phase")
	}
	session.UserStatus.TutorialPhase = enum.TutorialPhaseFavoriteMember
}
