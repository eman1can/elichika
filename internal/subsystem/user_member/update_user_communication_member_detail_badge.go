package user_member

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
)

func UpdateUserCommunicationMemberDetailBadge(session *userdata.Session, badge client.UserCommunicationMemberDetailBadge) {
	session.UserModel.UserCommunicationMemberDetailBadgeById.Set(badge.MemberMasterId, badge)
}
