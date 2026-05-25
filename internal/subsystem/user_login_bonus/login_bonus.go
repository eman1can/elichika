package user_login_bonus

import (
	"elichika/internal/client"
	"elichika/internal/config"
	"elichika/internal/enum"
	"elichika/internal/gamedata"
	"elichika/internal/generic"
	"elichika/internal/subsystem/user_present"
	"elichika/internal/userdata"
	"elichika/internal/userdata/database"
	"fmt"
	"time"
)

func latestLoginBonusTime(timePoint time.Time) time.Time {
	year, month, day := timePoint.Date()
	res := time.Date(year, month, day, 0, 0, int(*config.Conf.LoginBonusSecond), 0, timePoint.Location())
	if res.After(timePoint) {
		res = res.AddDate(0, 0, -1)
	}
	return res
}

func nextLoginBonusTime(timePoint time.Time) time.Time {
	return latestLoginBonusTime(timePoint).AddDate(0, 0, 1)
}

func incrementLoginBonus(session *userdata.Session, loginBonus *gamedata.LoginBonus, userLoginBonus database.UserLoginBonus, allowRestart bool) bool {
	lastUnlocked := latestLoginBonusTime(session.Time)
	if userLoginBonus.LastReceivedAt >= lastUnlocked.Unix() {
		return true
	}

	userLoginBonus.LastReceivedAt = session.Time.Unix()
	userLoginBonus.LastReceivedReward++

	if userLoginBonus.LastReceivedReward == loginBonus.LoginBonusRewards.Size() {
		if allowRestart {
			userLoginBonus.LastReceivedReward = 0
		} else {
			return true
		}
	}

	return false
}

func awardLoginBonusItems(session *userdata.Session, loginBonus *gamedata.LoginBonus, userLoginBonus database.UserLoginBonus, target *client.BootstrapLoginBonus) {
	naviLoginBonus := loginBonus.NaviLoginBonus()
	for i := range naviLoginBonus.LoginBonusRewards.Slice {
		if i < userLoginBonus.LastReceivedReward {
			naviLoginBonus.LoginBonusRewards.Slice[i].Status = enum.LoginBonusReceiveStatusReceived
		} else if i > userLoginBonus.LastReceivedReward {
			naviLoginBonus.LoginBonusRewards.Slice[i].Status = enum.LoginBonusReceiveStatusUnreceived
		} else {
			naviLoginBonus.LoginBonusRewards.Slice[i].Status = enum.LoginBonusReceiveStatusReceiving
		}
	}
	target.LoginBonuses.Append(naviLoginBonus)
	for _, content := range loginBonus.LoginBonusRewards.Slice[userLoginBonus.LastReceivedReward].LoginBonusContents.Slice {
		user_present.AddPresent(session, client.PresentItem{
			Content:          content,
			PresentRouteType: enum.PresentRouteTypeLoginBonus,
			PresentRouteId:   generic.NewNullable(loginBonus.LoginBonusId),
			ParamServer: generic.NewNullable(client.LocalizedText{
				DotUnderText: loginBonus.DotUnderText,
			}),
			ParamClient: generic.NewNullable(fmt.Sprint(userLoginBonus.LastReceivedReward + 1)),
		})
	}
	updateUserLoginBonus(session, userLoginBonus)
}
