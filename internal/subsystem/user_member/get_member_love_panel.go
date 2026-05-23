package user_member

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
)

func GetMemberLovePanel(session *userdata.Session, memberId int32) client.MemberLovePanel {
	panel, exist := session.MemberLovePanelDiffs[memberId]
	if exist {
		return panel
	}
	return GetOtherUserMemberLovePanel(session, int32(session.UserId), memberId)
}
