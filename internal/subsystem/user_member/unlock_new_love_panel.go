package user_member

import (
	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/generic"
	"elichika/internal/subsystem/user_info_trigger"
	"elichika/internal/userdata"
)

func UnlockNewLovePanel(session *userdata.Session, memberId, oldLoveLevel, newLoveLevel int32) {
	currentLovePanel := GetCurrentMemberLovePanel(session, memberId, oldLoveLevel)

	if currentLovePanel.AllUnlocked() && !currentLovePanel.IsLastPanel {
		nextPanel := session.Gamedata.MemberLovePanel[currentLovePanel.PanelId].NextPanel

		if nextPanel.LoveLevel <= newLoveLevel && nextPanel.LoveLevel > oldLoveLevel {
			user_info_trigger.AddTriggerBasic(session, client.UserInfoTriggerBasic{
				InfoTriggerType: enum.InfoTriggerTypeMemberLovePanelNew,
				ParamInt:        generic.NewNullable(nextPanel.Id)})
		}
	}
}
