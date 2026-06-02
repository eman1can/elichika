package gamedata

import (
	"log"

	"elichika/internal/utils"

	"xorm.io/xorm"
)

type StoryEventHistory struct {
	// from m_story_event_history_detail
	StoryEventId        int32  `xorm:"pk 'story_event_id'"`
	EventMasterId       int32  `xorm:"event_master_id"`
	BannerThumbnailPath string `xorm:"banner_thumbnail_path"`
	DetailThumbnailPath string `xorm:"detail_thumbnail_path"`
}

func loadStoryEventHistory(gamedata *Gamedata) {
	log.Println("Loading StoryEventHistory")
	gamedata.StoryEventHistory = make(map[int32]*StoryEventHistory)
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_story_event_history_detail").Find(&gamedata.StoryEventHistory)
	})
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadStoryEventHistory)
}
