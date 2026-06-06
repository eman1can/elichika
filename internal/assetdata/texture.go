package assetdata

import (
	"elichika/internal/utils"

	"xorm.io/xorm"
)

type Texture struct {
	AssetPath string `xorm:"pk 'asset_path'"`
	PackName  string `xorm:"pack_name"`
	Head      int    `xorm:"head"`
	Size      int32  `xorm:"size"`
	Key1      uint32 `xorm:"key1"`
	Key2      uint32 `xorm:"key2"`
}

func loadTexture(session *xorm.Session, ad *Assetdata) {
	var textures []*Texture

	err := session.Table("texture").Find(&textures)
	utils.CheckErr(err)

	for _, texture := range textures {
		ad.TextureByAssetPath[texture.AssetPath] = texture
	}
}
