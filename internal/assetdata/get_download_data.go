//go:build !dev

package assetdata

import (
	"fmt"
	"log"
)

type DownloadData struct {
	Locale       string
	File         string
	Package      string
	IsEntireFile bool // if this is set, the following fields are 0
	Start        int
	Size         int
}

func GetDownloadData(packname string) DownloadData {
	_, exist := Metapack[packname]
	if exist {
		return DownloadData{
			Locale:       NameToLocale[packname],
			File:         packname,
			Package:      fmt.Sprintf("meta%c", packname[0]),
			IsEntireFile: true,
		}
	}
	pack, exist := Pack[packname]
	if !exist {
		log.Panic(fmt.Sprint("package doesn't exist: ", packname))
	}
	if pack.Metapack == nil {
		return DownloadData{
			Locale:       NameToLocale[packname],
			File:         packname,
			Package:      fmt.Sprintf("pkg%c", packname[0]),
			IsEntireFile: true,
		}
	}

	return DownloadData{
		Locale:       NameToLocale[pack.Metapack.MetapackName],
		File:         pack.Metapack.MetapackName,
		Package:      fmt.Sprintf("meta%c", pack.Metapack.MetapackName[0]),
		IsEntireFile: false,
		Start:        pack.MetapackOffset,
		Size:         pack.FileSize,
	}
}
