package assetdata

import (
	"elichika/internal/utils"

	"xorm.io/xorm"
)

var SoundBySheetName = map[string]*Sound{}

type Sound struct {
	SheetName   string `xorm:"sheet_name"`
	AcbPackName string `xorm:"acb_pack_name"`
	AwbPackName string `xorm:"awb_pack_name"`
}

func loadSound(session *xorm.Session) {
	var sounds []*Sound

	err := session.Table("m_asset_sound").Find(&sounds)
	utils.CheckErr(err)

	for _, sound := range sounds {
		SoundBySheetName[sound.SheetName] = sound
	}
}
