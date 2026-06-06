package user_story_event_history

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
)

func IsEventStoryFinished(session *userdata.Session, storyEventId int32) bool {
	if _, ok := session.UserModel.UserStoryEventHistoryById.Get(storyEventId); ok {
		return ok
	}

	userStoryEventHistory := client.UserStoryEventHistory{
		StoryEventId: storyEventId,
	}

	return userdata.GenericDatabaseExist(session, "u_story_event_history", userStoryEventHistory)
}
