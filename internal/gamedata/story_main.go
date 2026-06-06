package gamedata

import (
	"log"

	"elichika/internal/utils"

	"xorm.io/xorm"
)

type StoryMainChapter struct {
	// from m_story_main_chapter
	Id                 int32  `xorm:"pk 'id'" json:"id"`
	Title              string `xorm:"title"`
	Description        string `xorm:"description"`
	ThumbnailAssetPath string `xorm:"thumbnail_asset_path"`

	// Linked from m_story_main_cell
	Cells      []int32 `xorm:"-" json:"-"`
	LastCellId int32   `xorm:"-" json:"-"`
}

func (smc *StoryMainChapter) populate(gamedata *Gamedata) {
	for _, cell := range gamedata.StoryMainChapterCell {
		if cell.ChapterId != smc.Id {
			continue
		}

		smc.Cells = append(smc.Cells, cell.Id)
		if smc.LastCellId < cell.Id {
			smc.LastCellId = cell.Id
		}
	}
}

func loadStoryMain(gamedata *Gamedata) {
	gamedata.StoryMainChapter = make(map[int32]*StoryMainChapter)
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_story_main_chapter").Find(&gamedata.StoryMainChapter)
	})
	utils.CheckErr(err)
	for _, storyMainChapter := range gamedata.StoryMainChapter {
		storyMainChapter.populate(gamedata)
	}

	log.Println("Loaded", len(gamedata.StoryMainChapter), "Main Story Chapters")
}

func init() {
	addLoadFunc(loadStoryMain)
	addPrequisite(loadStoryMain, loadStoryMainCell)
}
