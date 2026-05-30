package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"elichika/internal/config"
	"elichika/internal/utils"

	hwdecrypt "github.com/arina999999997/gohwdecrypt"
	"xorm.io/xorm"
)

type MasterNaviTimeline struct {
	Id                int64  `xorm:"id"`
	TimelineAssetPath string `xorm:"timeline_asset_path"`
}

type NaviTimeline struct {
	AssetPath string `xorm:"asset_path"`
	PackName  string `xorm:"pack_name"`
	Head      int64  `xorm:"head"`
	Size      int64  `xorm:"size"`
	Key1      uint32 `xorm:"key1"`
	Key2      uint32 `xorm:"key2"`
}

type AssetPackageMapping struct {
	PackName       string `xorm:"pack_name"`
	FileSize       int64  `xorm:"file_size"`
	MetaPackName   string `xorm:"metapack_name"`
	MetaPackOffset int64  `xorm:"metapack_offset"`
}

func main() {
	engine, err := xorm.NewEngine("sqlite", config.GlMasterdataPath+"masterdata.db")
	if err != nil {
		log.Fatal(err)
	}

	var masterTimelines []MasterNaviTimeline
	if err := engine.Table("m_navi_timeline").Find(&masterTimelines); err != nil {
		log.Fatal(err)
	}
	engine.Close()

	var idByAssetPath = make(map[string]int64)
	for _, timeline := range masterTimelines {
		idByAssetPath[timeline.TimelineAssetPath] = timeline.Id
	}

	engine, err = xorm.NewEngine("sqlite", config.GlMasterdataPath+"asset_a_en.db")
	if err != nil {
		log.Fatal(err)
	}

	var timelines []NaviTimeline
	if err := engine.Table("navi_timeline").Find(&timelines); err != nil {
		log.Fatal(err)
	}

	var packs []AssetPackageMapping
	if err := engine.Table("m_asset_package_mapping").Find(&packs); err != nil {
		log.Fatal(err)
	}

	engine.Close()

	var metapackByPackName = make(map[string]AssetPackageMapping)
	for _, pack := range packs {
		metapackByPackName[pack.PackName] = pack
	}

	for _, timeline := range timelines {
		id := idByAssetPath[timeline.AssetPath]

		offset := timeline.Head
		path := filepath.Join(config.StaticDataPath, "packs", timeline.PackName)

		if metapack, ok := metapackByPackName[timeline.PackName]; ok {
			if metapack.MetaPackName != "" {
				offset += metapack.MetaPackOffset
				path = filepath.Join(config.StaticDataPath, "packs", metapack.MetaPackName)
			}
		}

		file, err := os.Open(path)
		utils.CheckErr(err)
		output := make([]byte, timeline.Size)
		_, err = file.ReadAt(output, offset)
		utils.CheckErr(err)
		hwdecrypt.DecryptBuffer(&hwdecrypt.HwdKeyset{
			Key1: timeline.Key1,
			Key2: timeline.Key2,
			Key3: 12345,
		}, output)
		file.Close()

		os.WriteFile(filepath.Join(config.StaticDataPath, "timelines", fmt.Sprintf("%d.unityfs", id)), output, 0666)
	}
}
