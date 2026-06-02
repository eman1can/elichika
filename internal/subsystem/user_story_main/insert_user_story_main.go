package user_story_main

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
)

func InsertUserStoryMain(session *userdata.Session, storyMainMasterId int32) bool {
	userStoryMain := client.UserStoryMain{
		StoryMainMasterId: storyMainMasterId,
	}
	if userdata.GenericDatabaseExist(session, "u_story_main", userStoryMain) {
		return false
	}
	session.UserModel.UserStoryMainByStoryMainId.Set(storyMainMasterId, userStoryMain)
	// Before EOS, main story would also unlock scenes, but that part of the tutorial has been removed from the client
	return true
}
