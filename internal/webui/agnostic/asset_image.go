package agnostic

import (
	"errors"
	"os"

	"elichika/internal/assetdata"
	"elichika/internal/config"
	"elichika/internal/utils"

	hwdecrypt "github.com/arina999999997/gohwdecrypt"
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
	texture, exists := assetdata.TextureByAssetPath[assetPath]
	if !exists {
		return nil, errors.New("asset not found in database")
	}

	data := assetdata.GetDownloadData(texture.PackName)
	path := config.StaticDataPath + "packs/" + data.File
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	output := make([]byte, texture.Size)
	file, err := os.Open(path)
	utils.CheckErr(err)
	defer file.Close()
	_, err = file.ReadAt(output, int64(data.Start+texture.Head))
	utils.CheckErr(err)
	hwdecrypt.DecryptBuffer(&hwdecrypt.HwdKeyset{
		Key1: texture.Key1,
		Key2: texture.Key2,
		Key3: 12345,
	}, output)

	return output, nil
}
