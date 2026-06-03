package user_story_event_history

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
)

func GetEventStory(session *userdata.Session, storyEventId int32) *client.UserStoryEventHistory {
	if story, ok := session.UserModel.UserStoryEventHistoryById.Get(storyEventId); ok {
		return story
	}

	return &client.UserStoryEventHistory{
		StoryEventId: storyEventId,
		IsNew:        true,
	}
}