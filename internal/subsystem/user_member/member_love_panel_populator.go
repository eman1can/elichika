package user_member

import (
	"elichika/internal/userdata"
)

func memberLovePanelPopulator(session *userdata.Session) {
	for _, masterPanel := range session.Gamedata.MemberLovePanel {
		memberId := *masterPanel.MemberMasterId
		session.MemberLovePanels[memberId][masterPanel.Id] = GetMemberLovePanel(session, memberId, masterPanel.Id)
	}
	session.MemberLovePanelList = GetLovePanelList(session)
}

func init() {
	userdata.AddContentPopulator(memberLovePanelPopulator)
}
