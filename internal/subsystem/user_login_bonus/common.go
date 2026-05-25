package user_login_bonus

import (
	"elichika/internal/client"
	"elichika/internal/gamedata"
	"elichika/internal/userdata"
)

type HandlerType = func(string, *userdata.Session, *gamedata.LoginBonus, *client.BootstrapLoginBonus)

var handler = map[string]HandlerType{}

func init() {
	handler["limited_login_bonus"] = limitedLoginBonusHandler
	handler["normal_login_bonus"] = normalLoginBonusHandler
	handler["birthday_login_bonus"] = birthdayLoginBonusHandler
}
