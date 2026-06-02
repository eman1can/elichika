package user_member

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

func GetOtherUserMemberLovePanel(session *userdata.Session, userId, memberId int32, panelId int32) client.MemberLovePanel {
	isLastPanel := session.Gamedata.MemberLovePanel[panelId].NextPanel == nil

	result := client.MemberLovePanel{
		IsLastPanel: isLastPanel,
	}
	exist, err := session.Db.Table("u_member_love_panel").
		Where("user_id = ? AND member_id = ? AND panel_id = ?", userId, memberId, panelId).
		Get(&result)
	utils.CheckErr(err)

	if !exist {
		return client.MemberLovePanel{
			MemberId:    memberId,
			PanelId:     panelId,
			Status:      0,
			IsLastPanel: isLastPanel,
		}
	}

	return result
}
