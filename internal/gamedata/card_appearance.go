package gamedata

import (
	"elichika/internal/utils"
	"log"

	"xorm.io/xorm"
)

type CardAppearance struct {
	// from m_card_appearance
	CardMasterId   int32 `xorm:"pk 'card_m_id'"`
	AppearanceType int32 `xorm:"'appearance_type'"`

	CardName      string `xorm:"'card_name'"`
	Pronunciation string `xorm:"'pronunciation'"`

	ImageAssetPath          string `xorm:"'image_asset_path'"`
	ThumbnailAssetPath      string `xorm:"'thumbnail_asset_path'"`
	StillThumbnailAssetPath string `xorm:"'still_thumbnail_asset_path'"`
	BackgroundAssetPath     string `xorm:"'background_asset_path'"`
	LiveDeckAssetPathId     string `xorm:"'live_deck_asset_path_id'"`
}

func (cardAppearance *CardAppearance) populate(gamedata *Gamedata) {
	if cardAppearance.AppearanceType == 1 {
		gamedata.CardAppearance[cardAppearance.CardMasterId] = cardAppearance
	} else {
		gamedata.CardIdolAppearance[cardAppearance.CardMasterId] = cardAppearance
	}
}

func loadCardAppearance(gamedata *Gamedata) {
	log.Println("Loading CardAppearance")

	var appearances []CardAppearance
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_card_appearance").Find(&appearances)
	})
	utils.CheckErr(err)

	gamedata.CardAppearance = make(map[int32]*CardAppearance)
	gamedata.CardIdolAppearance = make(map[int32]*CardAppearance)
	for _, cardAppearance := range appearances {
		cardAppearance.populate(gamedata)
	}
}

func init() {
	addLoadFunc(loadCardAppearance)
}
