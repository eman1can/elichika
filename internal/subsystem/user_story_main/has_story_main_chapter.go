package user_story_main

import (
	"elichika/internal/userdata"
)

func HasStoryMainChapter(session *userdata.Session, storyMainMasterId int32) bool {
	all := true

	if story, ok := session.Gamedata.StoryMainChapter[storyMainMasterId]; ok {
		for _, cellId := range story.Cells {
			all = all && HasStoryMainCell(session, cellId)
		}
	}

	return all
}
