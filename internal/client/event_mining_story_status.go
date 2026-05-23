package client

import "elichika/internal/generic"

type EventMiningStoryStatus struct {
	ReadStoryNumber int32                          `json:"read_story_number"`
	Stories         generic.List[EventMiningStory] `json:"stories"`
}
