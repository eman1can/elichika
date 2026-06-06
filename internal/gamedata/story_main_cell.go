package gamedata

import (
	"log"

	"elichika/internal/utils"

	"xorm.io/xorm"
)

type StoryMainChapterCell struct {
	Id                 int32   `xorm:"pk 'id'"`
	ChapterId          int32   `xorm:"chapter_id"`
	DisplayOrder       int32   `xorm:"display_order"`
	Title              string  `xorm:"title"`
	Description        string  `xorm:"summary"`
	ThumbnailAssetPath *string `xorm:"thumbnail_asset_path"`
	LiveDifficultyId   *int32  `xorm:"live_difficulty_id"`
}

func loadStoryMainCell(gamedata *Gamedata) {
	gamedata.StoryMainChapterCell = make(map[int32]*StoryMainChapterCell)
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_story_main_cell").Find(&gamedata.StoryMainChapterCell)
	})
	utils.CheckErr(err)

	log.Println("Loaded", len(gamedata.StoryMainChapterCell), "Main Story Chapter Cells")
}

func init() {
	addLoadFunc(loadStoryMainCell)
}
