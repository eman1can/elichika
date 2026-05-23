package user_story_linkage

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
)

func InsertUserStoryLinkage(session *userdata.Session, storyLinkageCellMasterId int32) {
	userStoryLinkage := client.UserStoryLinkage{
		StoryLinkageCellMasterId: storyLinkageCellMasterId,
	}
	if !userdata.GenericDatabaseExist(session, "u_story_linkage", userStoryLinkage) {
		session.UserModel.UserStoryLinkageById.Set(storyLinkageCellMasterId, userStoryLinkage)
	}
}
