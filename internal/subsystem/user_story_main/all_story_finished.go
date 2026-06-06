package user_story_main

import "elichika/internal/userdata"

func AllStoryFinished(session *userdata.Session) bool {
	all := true

	for _, story := range session.Gamedata.StoryMainChapter {
		all = all && HasStoryMainChapter(session, story.Id)
	}

	return all
}
