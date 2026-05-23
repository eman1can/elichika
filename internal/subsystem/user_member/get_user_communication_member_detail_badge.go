package user_member

import (
	"log"

	"elichika/internal/client"
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

func GetUserCommunicationMemberDetailBadge(session *userdata.Session, memberMasterId int32) client.UserCommunicationMemberDetailBadge {
	ptr, exist := session.UserModel.UserCommunicationMemberDetailBadgeById.Get(memberMasterId)
	if exist {
		return *ptr
	}
	badge := client.UserCommunicationMemberDetailBadge{}
	exist, err := session.Db.Table("u_communication_member_detail_badge").
		Where("user_id = ? AND member_master_id = ?", session.UserId, memberMasterId).Get(&badge)
	utils.CheckErr(err)
	if !exist {
		// always inserted at login if not exist
		log.Panic("member not found")
	}
	return badge
}
