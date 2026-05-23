package user_member_guild

import (
	"time"

	"elichika/internal/userdata"
)

// the year is based on the end of the period, not beginning
func GetMemberGuildIdYear(session *userdata.Session, memberGuildId int32) int32 {
	endTime := session.Gamedata.MemberGuildPeriod.StartAt + session.Gamedata.MemberGuildPeriod.OneCycleSecs*int64(memberGuildId)
	return int32(time.Unix(endTime, 0).Year())
}
