package user_story_member

import (
	"elichika/internal/subsystem/user_member"
	"elichika/internal/userdata"
)

func AllStoryMembersUnlocked(session *userdata.Session, memberId int32) bool {
	all := true
	for _, masterStoryMember := range session.Gamedata.StoryMember {
		if masterStoryMember.MemberMId != memberId {
			continue
		}

		member := user_member.GetMember(session, memberId)
		if member.LoveLevel < masterStoryMember.LoveLevel {
			continue
		}

		memberStory := GetStoryMember(session, masterStoryMember.Id)
		all = all && !memberStory.IsNew
	}
	return all
}
