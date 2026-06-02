package user_story_member

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

func GetStoryMember(session *userdata.Session, storyMemberMasterId int32) client.UserStoryMember {
	if ptr, exist := session.UserModel.UserStoryMemberById.Get(storyMemberMasterId); exist {
		return *ptr
	}

	result := client.UserStoryMember{}
	exist, err := session.Db.Table("u_story_member").
		Where("user_id = ? AND story_member_master_id = ?", session.UserId, storyMemberMasterId).
		Get(&result)
	utils.CheckErr(err)

	if !exist {
		return client.UserStoryMember{
			StoryMemberMasterId: storyMemberMasterId,
			IsNew:               true,
			AcquiredAt:          session.Time.Unix(),
		}
	}

	return result
}
