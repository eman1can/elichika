package user_story_member

import (
	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/subsystem/user_content"
	"elichika/internal/userdata"
)

func storyMemberHandler(session *userdata.Session, content *client.Content) any {
	InsertMemberStory(session, content.ContentId)
	content.ContentAmount = 0
	return nil
}

func init() {
	user_content.AddContentHandler(enum.ContentTypeStoryMember, storyMemberHandler)
}
