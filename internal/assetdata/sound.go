package assetdata

import (
	"elichika/internal/utils"

	"xorm.io/xorm"
)

type Sound struct {
	SheetName   string `xorm:"sheet_name"`
	AcbPackName string `xorm:"acb_pack_name"`
	AwbPackName string `xorm:"awb_pack_name"`
}

func loadSound(session *xorm.Session, ad *Assetdata) {
	var sounds []*Sound

	err := session.Table("m_asset_sound").Find(&sounds)
	utils.CheckErr(err)

	for _, sound := range sounds {
		ad.SoundBySheetName[sound.SheetName] = sound
	}
}
