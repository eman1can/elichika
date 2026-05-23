package user_member

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
)

func UpdateMemberLovePanel(session *userdata.Session, panel client.MemberLovePanel) {
	session.MemberLovePanelDiffs[panel.MemberId] = panel
}
