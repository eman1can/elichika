package user_login_bonus

import (
	"log"

	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/gamedata"
	"elichika/internal/userdata"
)

func limitedLoginBonusHandler(_ string, session *userdata.Session, loginBonus *gamedata.LoginBonus, target *client.BootstrapLoginBonus) {
	if loginBonus.LoginBonusType != enum.LoginBonusTypeNormal {
		log.Panic("wrong handler used")
	}
	userLoginBonus := getUserLoginBonus(session, loginBonus.LoginBonusId)
	if incrementLoginBonus(session, loginBonus, userLoginBonus, false) {
		return
	}
	awardLoginBonusItems(session, loginBonus, userLoginBonus, target)
	updateUserLoginBonus(session, userLoginBonus)
}
