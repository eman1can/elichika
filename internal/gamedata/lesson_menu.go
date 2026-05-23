package gamedata

import (
	"log"

	"elichika/internal/client"
	"elichika/internal/generic/drop"
	"elichika/internal/utils"

	"xorm.io/xorm"
)

type LessonMenu struct {
	Id int32 `xorm:"pk"`

	DefaultItemDrop *drop.WeightedDropList[client.LessonDropItem]           `xorm:"-"`
	ItemDrop        map[int32]*drop.WeightedDropList[client.LessonDropItem] `xorm:"-"`
}

func (lm *LessonMenu) populate(gamedata *Gamedata) {
	type LessonDropContent struct {
		ContentType   int32
		ContentId     int32
		ContentAmount int32
		Weight        int32
		Rarity        int32
	}

	var contents []LessonDropContent
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_lesson_drop_content").Where("lesson_menu_master_id == ?", lm.Id).Find(&contents)
	})
	utils.CheckErr(err)
	lm.DefaultItemDrop = &drop.WeightedDropList[client.LessonDropItem]{}
	for _, content := range contents {
		lm.DefaultItemDrop.AddItem(client.LessonDropItem{
			ContentType:   content.ContentType,
			ContentId:     content.ContentId,
			ContentAmount: content.ContentAmount,
			DropRarity:    content.Rarity,
		}, content.Weight)
	}

	type LessonEnhancingItemDropRate struct {
		LessonEnhancingItemId int32
		TargetRarity          int32
		MagnificationWeight   int32
	}
	var enhancingItems []LessonEnhancingItemDropRate
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_lesson_enhancing_item_effect_drop_rate").Find(&enhancingItems)
	})
	utils.CheckErr(err)
	lm.ItemDrop = map[int32]*drop.WeightedDropList[client.LessonDropItem]{}
	for _, rate := range enhancingItems {
		if lm.ItemDrop[rate.LessonEnhancingItemId] == nil {
			lm.ItemDrop[rate.LessonEnhancingItemId] = &drop.WeightedDropList[client.LessonDropItem]{}
		}
		for _, content := range contents {
			if content.Rarity == rate.TargetRarity {
				lm.ItemDrop[rate.LessonEnhancingItemId].AddItem(client.LessonDropItem{
					ContentType:   content.ContentType,
					ContentId:     content.ContentId,
					ContentAmount: content.ContentAmount,
					DropRarity:    content.Rarity,
				}, content.Weight*rate.MagnificationWeight/10000)
			}
		}
	}
}

func loadLessonMenu(gamedata *Gamedata) {
	log.Println("Loading LessonMenu")
	gamedata.LessonMenu = make(map[int32]*LessonMenu)
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_lesson_menu").Find(&gamedata.LessonMenu)
	})
	utils.CheckErr(err)
	for _, lessonMenu := range gamedata.LessonMenu {
		lessonMenu.populate(gamedata)
	}
}

func init() {
	addLoadFunc(loadLessonMenu)
}
