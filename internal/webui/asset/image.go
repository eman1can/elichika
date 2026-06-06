package asset

import (
	"errors"
	"os"

	"elichika/internal/assetdata"
	"elichika/internal/config"
	"elichika/internal/utils"

	hwdecrypt "github.com/arina999999997/gohwdecrypt"
	"github.com/gin-gonic/gin"
)

func LoadAssetImage(ctx *gin.Context, assetPath string) ([]byte, error) {
	ad := ctx.MustGet("assetdata").(*assetdata.Assetdata)
	texture, exists := ad.TextureByAssetPath[assetPath]
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
