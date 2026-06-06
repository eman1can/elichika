package user_story_main

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
)

func HasStoryMainCell(session *userdata.Session, storyMainCellId int32) bool {
	userStoryMain := client.UserStoryMain{
		StoryMainCellId: storyMainCellId,
	}
	return userdata.GenericDatabaseExist(session, "u_story_main", userStoryMain)
}
