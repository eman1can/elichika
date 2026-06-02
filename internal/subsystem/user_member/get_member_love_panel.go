package user_member

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
)

func GetOtherUserLovePanelMemberList(session *userdata.Session, userId int32, memberId int32) client.MemberLovePanelList {
	panelInfo := client.MemberLovePanelList{MemberId: memberId}
	for panelId, masterPanel := range session.Gamedata.MemberLovePanel {
		if *masterPanel.MemberMasterId != memberId {
			continue
		}

		panel := GetOtherUserMemberLovePanel(session, userId, memberId, panelId)
		if panel.NoneUnlocked() {
			continue
		}

		for ix, cellId := range masterPanel.CellIds {
			if panel.Unlocked(ix) {
				panelInfo.MemberLovePanelCellIds.Append(cellId)
			}
		}
	}
	return panelInfo
}

func GetLovePanelMemberList(session *userdata.Session, memberId int32) client.MemberLovePanelList {
	return GetOtherUserLovePanelMemberList(session, session.UserId, memberId)
}

// GetLovePanelList
// Get the list of unlocked love panel cells grouped by member
// Requires that session.MemberLovePanels has been appropriately populated
func GetLovePanelList(session *userdata.Session) []client.MemberLovePanelList {
	var resp []client.MemberLovePanelList

	for memberId := range session.Gamedata.Member {
		resp = append(resp, GetLovePanelMemberList(session, memberId))
	}

	return resp
}

func GetMemberLovePanel(session *userdata.Session, memberId int32, panelId int32) client.MemberLovePanel {
	if panel, exist := session.MemberLovePanelDiffs[memberId][panelId]; exist {
		return panel
	}
	return GetOtherUserMemberLovePanel(session, session.UserId, memberId, panelId)
}

func GetCurrentMemberLovePanel(session *userdata.Session, memberId int32, loveLevel int32) client.MemberLovePanel {
	masterPanel := session.Gamedata.MemberFirstLovePanel[memberId]
	for masterPanel.NextPanel != nil {
		panel := GetMemberLovePanel(session, memberId, masterPanel.Id)

		if panel.AllUnlocked() && loveLevel >= masterPanel.LoveLevel {
			masterPanel = masterPanel.NextPanel
			continue
		}

		return panel
	}

	return GetMemberLovePanel(session, memberId, masterPanel.Id)
}
