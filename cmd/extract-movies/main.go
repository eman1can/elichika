package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"elichika/internal/config"

	"github.com/schollz/progressbar/v3"
	"xorm.io/xorm"
)

type Movie struct {
	Pavement string `xorm:"pavement"`
	PackName string `xorm:"pack_name"`
}

type AssetPackageMapping struct {
	PackName       string `xorm:"pack_name"`
	FileSize       int64  `xorm:"file_size"`
	MetaPackName   string `xorm:"metapack_name"`
	MetaPackOffset int64  `xorm:"metapack_offset"`
}

func loadPackBytes(packs *[]AssetPackageMapping, packName string) ([]byte, error) {
	for _, pack := range *packs {
		if pack.PackName == packName {
			var err error
			var data []byte
			var offset int64
			if pack.MetaPackName != "" {
				data, err = os.ReadFile(filepath.Join(config.StaticDataPath, "packs", pack.MetaPackName))
				offset = pack.MetaPackOffset
			} else {
				data, err = os.ReadFile(filepath.Join(config.StaticDataPath, "packs", packName))
				offset = int64(0)
			}

			if err != nil {
				return []byte{}, err
			}

			size := pack.FileSize
			return data[offset : offset+size], nil
		}
	}

	return []byte{}, fmt.Errorf("pack %s not found", packName)
}

func main() {
	engine, err := xorm.NewEngine("sqlite", config.GlMasterdataPath+"asset_a_en.db")
	if err != nil {
		log.Fatal(err)
	}

	var movies []Movie
	if err := engine.Table("m_movie").Find(&movies); err != nil {
		log.Fatal(err)
	}

	var packs []AssetPackageMapping
	if err := engine.Table("m_asset_package_mapping").Find(&packs); err != nil {
		log.Fatal(err)
	}

	err = engine.Close()
	if err != nil {
		log.Fatal(err)
	}

	bar := progressbar.Default(int64(len(movies)), "Extracting movies")

	var wg sync.WaitGroup
	sem := make(chan struct{}, 20)
	for _, movie := range movies {
		wg.Add(1)
		sem <- struct{}{} // acquire slot

		go func(movie Movie) {
			defer wg.Done()
			defer bar.Add(1)
			defer func() { <-sem }() // release slot

			usmPath := filepath.Join(config.StaticDataPath, "sounds", "usm", movie.PackName+".usm")

			if _, err := os.Stat(usmPath); err != nil {
				usmData, err := loadPackBytes(&packs, movie.PackName)
				if err != nil {
					log.Fatal(err)
				}

				err = os.WriteFile(usmPath, usmData, 0666)
				if err != nil {
					log.Fatal(err)
				}
			}
		}(movie)
	}

	wg.Wait()
}
