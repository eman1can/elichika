package user_story_main

import "elichika/internal/userdata"

func AllStoryFinished(session *userdata.Session) bool {
	all := true

	for _, story := range session.Gamedata.StoryMainChapter {
		for _, cell := range story.Cells {
			all = all && IsStoryFinished(session, cell)
		}
	}

	return all
}
