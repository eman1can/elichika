package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"elichika/internal/config"

	"github.com/eman1can/sound_decrypt/awb"
	"github.com/eman1can/sound_decrypt/wav"
	"github.com/schollz/progressbar/v3"
	"xorm.io/xorm"
)

type SoundSheet struct {
	SheetName   string `xorm:"sheet_name"`
	AcbPackName string `xorm:"acb_pack_name"`
	AwbPackName string `xorm:"awb_pack_name"`
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

	var sheets []SoundSheet
	if err := engine.Table("m_asset_sound").Find(&sheets); err != nil {
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

	bar := progressbar.Default(int64(len(sheets)), "Extracting sounds")

	var wg sync.WaitGroup
	sem := make(chan struct{}, 20)
	for _, sheet := range sheets {
		wg.Add(1)
		sem <- struct{}{} // acquire slot

		go func(sheet SoundSheet) {
			defer wg.Done()
			defer bar.Add(1)
			defer func() { <-sem }() // release slot

			wavPath := filepath.Join(config.StaticDataPath, "sounds", "wav", sheet.SheetName+".wav")
			awbPath := filepath.Join(config.StaticDataPath, "sounds", "awb", sheet.SheetName+".awb")
			acbPath := filepath.Join(config.StaticDataPath, "sounds", "acb", sheet.SheetName+".acb")

			if _, err := os.Stat(acbPath); err != nil {
				if sheet.AcbPackName != "" {
					acbData, err := loadPackBytes(&packs, sheet.AcbPackName)
					if err != nil {
						log.Fatal(err)
					}

					err = os.WriteFile(acbPath, acbData, 0666)
					if err != nil {
						log.Fatal(err)
					}
				}
			}

			if sheet.AwbPackName != "" {
				awbData, err := loadPackBytes(&packs, sheet.AwbPackName)
				if err != nil {
					log.Println(err)
					return
				}

				if _, err := os.Stat(awbPath); err != nil {
					err = os.WriteFile(awbPath, awbData, 0666)
					if err != nil {
						log.Fatal(err)
					}
				}

				if _, err := os.Stat(wavPath); err != nil {
					awbFile, err := awb.LoadAWB(awbData, 6498535309877346413)
					if err != nil {
						log.Println(err)
						return
					}

					if awbFile.TotalSubsongs > 1 {
						fmt.Println("invalid number of subsongs")
						return
					}

					for _, hcaFile := range awbFile.Subfiles {
						buf := new(bytes.Buffer)

						if err := wav.WriteWAV(hcaFile, buf); err != nil {
							fmt.Fprintf(os.Stderr, "error writing %s: %v\n", wavPath, err)
							continue
						}

						data := buf.Bytes()
						err = os.WriteFile(wavPath, data, 0777)
						if err != nil {
							fmt.Fprintf(os.Stderr, "error writing %s: %v\n", wavPath, err)
						}
						break
					}
				}
			}
		}(sheet)
	}

	wg.Wait()
}
