package user_member

import (
	"slices"

	"elichika/internal/client"
	"elichika/internal/config"
	"elichika/internal/enum"
	"elichika/internal/generic"
	"elichika/internal/subsystem/user_content"
	"elichika/internal/subsystem/user_info_trigger"
	"elichika/internal/userdata"
)

func UnlockMemberLovePanelCells(session *userdata.Session, memberId int32, panelId int32, lovePanelCellIds []int32) client.MemberLovePanelList {
	masterLovePanel := session.Gamedata.MemberLovePanel[panelId]
	panel := GetMemberLovePanel(session, memberId, panelId)

	for ix, cellId := range masterLovePanel.CellIds {
		if slices.Contains(lovePanelCellIds, cellId) {
			panel.Status |= 1 << ix

			for _, resource := range session.Gamedata.MemberLovePanelCell[cellId].Resources {
				if config.Conf.ResourceConfig().ConsumePracticeItems {
					user_content.RemoveContent(session, resource)
				}
			}
		}
	}

	if panel.AllUnlocked() {
		member := GetMember(session, panel.MemberId)
		if (masterLovePanel.NextPanel != nil) && (masterLovePanel.NextPanel.LoveLevel <= member.LoveLevel) {
			user_info_trigger.AddTriggerBasic(session, client.UserInfoTriggerBasic{
				InfoTriggerType: enum.InfoTriggerTypeMemberLovePanelNew,
				ParamInt:        generic.NewNullable(masterLovePanel.NextPanel.Id)})
		}
	}

	UpdateMemberLovePanel(session, panel)
	return GetLovePanelMemberList(session, memberId)
}
