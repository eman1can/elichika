package user_story_member

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
)

func UpdateStoryMember(session *userdata.Session, userStoryMember client.UserStoryMember) {
	session.UserModel.UserStoryMemberById.Set(userStoryMember.StoryMemberMasterId, userStoryMember)
}
