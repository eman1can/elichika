package user_member

import (
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

func memberLovePanelsFinalizer(session *userdata.Session) {
	// Update MemberLovePanels with Updated Data and sync database with changes
	for _, memberPanels := range session.MemberLovePanelDiffs {
		for _, panel := range memberPanels {
			session.MemberLovePanels[panel.MemberId][panel.PanelId] = panel

			affected, err := session.Db.Table("u_member_love_panel").
				Where("user_id = ? AND member_id = ? AND panel_id = ?", session.UserId, panel.MemberId, panel.PanelId).
				AllCols().
				Update(panel)
			utils.CheckErr(err)
			if affected == 0 {
				userdata.GenericDatabaseInsert(session, "u_member_love_panel", panel)
			}
		}
	}
	session.MemberLovePanelList = GetLovePanelList(session)
}

func init() {
	userdata.AddFinalizer(memberLovePanelsFinalizer)
}
