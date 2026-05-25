package user_login_bonus

import (
	"log"

	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/gamedata"
	"elichika/internal/generic"
	"elichika/internal/item"
	"elichika/internal/subsystem/user_present"
	"elichika/internal/userdata"

	"fmt"
	"math/rand"
	"time"
)

func birthdayLoginBonusHandler(mode string, session *userdata.Session, loginBonus *gamedata.LoginBonus, target *client.BootstrapLoginBonus) {
	if loginBonus.LoginBonusType != enum.LoginBonusTypeBirthday {
		log.Panic("wrong handler used")
	}
	year, month, day := session.Time.Date()
	mmdd := int32(month)*100 + int32(day)
	list, exists := session.Gamedata.MemberByBirthday[mmdd]
	if !exists { // no one with this birthday
		return
	}
	userLoginBonus := getUserLoginBonus(session, loginBonus.LoginBonusId)
	lastUnlocked := time.Date(year, month, day, 0, 0, 0, 0, session.Time.Location())
	if userLoginBonus.LastReceivedAt >= lastUnlocked.Unix() { // already got it
		return
	}
	userLoginBonus.LastReceivedAt = session.Time.Unix()

	for _, member := range list {
		memorialCount := int32(50)
		if session.UserStatus.MemberGuildMemberMasterId.HasValue &&
			session.UserStatus.MemberGuildMemberMasterId.Value == member.Id {
			memorialCount += 25
		}

		loginRewards := []client.Content{
			item.MementoPiece(member.Id).Amount(2),
			item.MemorialPiece(member.Id).Amount(memorialCount),
			item.StarGem.Amount(50),
		}

		naviLoginBonus := loginBonus.NaviLoginBonus()
		naviLoginBonus.LoginBonusRewards.Append(
			client.LoginBonusRewards{
				Day:          1,
				Status:       enum.LoginBonusReceiveStatusReceiving,
				ContentGrade: generic.NewNullable(enum.LoginBonusContentGradeRare),
				LoginBonusContents: generic.Array[client.Content]{
					Slice: loginRewards,
				},
			},
		)

		for _, content := range loginRewards {
			user_present.AddPresent(session, client.PresentItem{
				Content:          content,
				PresentRouteType: enum.PresentRouteTypeLoginBonus,
				PresentRouteId:   generic.NewNullable(loginBonus.LoginBonusId),
				ParamServer: generic.NewNullable(client.LocalizedText{
					DotUnderText: fmt.Sprintf("k.m_dic_member_name_%d birthday_login_bonus", member.Id),
				}),
			})
		}

		// Choose the appropriate background and costume for the member
		memberLoginBonusBirthday := member.MemberLoginBonusBirthdays[0]
		switch mode {
		case "random":
			memberLoginBonusBirthday = member.MemberLoginBonusBirthdays[rand.Intn(len(member.MemberLoginBonusBirthdays))]
			break
		case "latest":
		case "equipped":
		default:
			log.Panic("not supported")
		}
		target.BirthdayMember.Append(client.LoginBonusBirthDayMember{
			MemberMasterId: generic.NewNullable(member.Id),
			SuitMasterId:   generic.NewNullable(memberLoginBonusBirthday.SuitMasterId),
		})
		naviLoginBonus.BackgroundId = memberLoginBonusBirthday.Id
		target.BirthdayLoginBonuses.Append(naviLoginBonus)
	}
	updateUserLoginBonus(session, userLoginBonus)
}
