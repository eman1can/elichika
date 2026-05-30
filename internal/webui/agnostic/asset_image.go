package agnostic

import (
	"elichika/internal/config"
	"elichika/internal/utils"
	"errors"
	"os"

	hwdecrypt "github.com/arina999999997/gohwdecrypt"
	"xorm.io/xorm"
)

type Texture struct {
	AssetPath string `xorm:"asset_path"`
	PackName  string `xorm:"pack_name"`
	Head      int    `xorm:"head"`
	Size      int    `xorm:"size"`
	Key1      uint32 `xorm:"key1"`
	Key2      uint32 `xorm:"key2"`
}

type PackageMapping struct {
	PackageKey     string  `xorm:"package_key"`
	PackName       string  `xorm:"pack_name"`
	FileSize       int     `xorm:"file_size"`
	MetapackName   *string `xorm:"metapack_name"`
	MetapackOffset int     `xorm:"metapack_offset"`
	Category       int     `xorm:"category"`
}

func loadAssetImage(assetPath string) ([]byte, error) {
	// TODO: Load this data into gamedata
	engine, err := xorm.NewEngine("sqlite", config.GlMasterdataPath+"asset_a_en.db")
	if err != nil {
		return nil, err
	}

	var textures []Texture
	err = engine.Table("texture").Where("asset_path = ?", assetPath).Limit(1).Find(&textures)
	if err != nil {
		return nil, err
	}
	if len(textures) == 0 {
		return nil, errors.New("no textures found")
	}

	var mappings []PackageMapping
	err = engine.Table("m_asset_package_mapping").Where("pack_name = ?", textures[0].PackName).Limit(1).Find(&mappings)
	if err != nil {
		return nil, err
	}
	if len(mappings) == 0 {
		return nil, errors.New("no packages found")
	}

	err = engine.Close()
	if err != nil {
		return nil, err
	}

	path := config.StaticDataPath + "packs/" + textures[0].PackName
	offset := textures[0].Head
	if mappings[0].MetapackName != nil {
		path = config.StaticDataPath + "packs/" + *mappings[0].MetapackName
		offset += mappings[0].MetapackOffset
	}

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	file, err := os.Open(path)
	utils.CheckErr(err)
	defer file.Close()
	output := make([]byte, textures[0].Size)
	_, err = file.ReadAt(output, int64(offset))
	utils.CheckErr(err)
	hwdecrypt.DecryptBuffer(&hwdecrypt.HwdKeyset{
		Key1: textures[0].Key1,
		Key2: textures[0].Key2,
		Key3: 12345,
	}, output)

	return output, nil
}
