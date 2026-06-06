package user_story_event_history

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
)

func UnlockEventStory(session *userdata.Session, storyEventId int32) {
	userStoryEventHistory := client.UserStoryEventHistory{
		StoryEventId: storyEventId,
	}
	if userdata.GenericDatabaseExist(session, "u_story_event_history", userStoryEventHistory) {
		return
	}
	session.UserModel.UserStoryEventHistoryById.Set(storyEventId, userStoryEventHistory)
}
