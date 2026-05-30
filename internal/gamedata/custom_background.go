package gamedata

import (
	"log"

	"elichika/internal/utils"

	"xorm.io/xorm"
)

type CustomBackground struct {
	// from m_custom_background
	Id                 int32  `xorm:"pk 'id'"`
	Name               string `xorm:"'name'"`
	ThumbnailAssetPath string `xorm:"'thumbnail_asset_path'"`
	DisplayOrder       int32  `xorm:"'display_order'"`
}

func loadCustomBackground(gamedata *Gamedata) {
	log.Println("Loading CustomBackground")
	gamedata.CustomBackground = make(map[int32]*CustomBackground)
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_custom_background").Find(&gamedata.CustomBackground)
	})
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadCustomBackground)
}
