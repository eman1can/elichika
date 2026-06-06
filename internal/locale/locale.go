package locale

import (
	"fmt"
	"log"
	"time"

	"elichika/internal/assetdata"
	"elichika/internal/config"
	"elichika/internal/db"
	"elichika/internal/gamedata"
	"elichika/internal/serverdata"
	"elichika/internal/utils"

	"xorm.io/xorm"
)

// create one engine for each potential file being read
// each locale is free to create and store its own session
var engines = map[string]*xorm.Engine{}

func GetEngine(path string) *xorm.Engine {
	engine, exist := engines[path]
	if exist {
		return engine
	}
	engine, err := xorm.NewEngine("sqlite", path)
	utils.CheckErr(err)
	engines[path] = engine
	return engine
}

type Locale struct {
	Path             string
	Language         string
	StartupKey       []byte
	MasterVersion    string
	Gamedata         *gamedata.Gamedata
	Dictionary       *gamedata.Dictionary
	AssetdataAndroid *assetdata.Assetdata
	AssetdataIos     *assetdata.Assetdata
}

func (locale *Locale) populate(syncChannel chan struct{}) {
	locale.Dictionary = new(gamedata.Dictionary)
	locale.Dictionary.Init(locale.Path, locale.Language)

	locale.Gamedata = new(gamedata.Gamedata)
	masterdataDb, err := db.NewDatabase(locale.Path + "masterdata.db")
	utils.CheckErr(err)
	locale.Gamedata.Init(locale.Language, masterdataDb, serverdata.Database, locale.Dictionary, syncChannel)
}

func (locale *Locale) loadAsset() {
	locale.AssetdataAndroid = new(assetdata.Assetdata)
	assetDbAndroid := GetEngine(fmt.Sprintf("%s/asset_a_%s.db", locale.Path, locale.Language))
	assetDbAndroid.SetMaxOpenConns(50)
	assetDbAndroid.SetMaxIdleConns(10)
	locale.AssetdataAndroid.Init(locale.Language, assetDbAndroid)

	locale.AssetdataIos = new(assetdata.Assetdata)
	assetDbIos := GetEngine(fmt.Sprintf("%s/asset_i_%s.db", locale.Path, locale.Language))
	assetDbIos.SetMaxOpenConns(50)
	assetDbIos.SetMaxIdleConns(10)
	locale.AssetdataIos.Init(locale.Language, assetDbIos)
}

var (
	Locales map[string]*Locale
)

func addLocale(path, language, masterVersion, startUpKey string) {
	locale := Locale{
		Path:          path,
		Language:      language,
		MasterVersion: masterVersion,
		StartupKey:    []byte(startUpKey),
	}
	Locales[language] = &locale
}

func init() {
	start := time.Now()
	gamedata.GenerateLoadOrder()
	Locales = make(map[string](*Locale))
	syncChannel := make(chan struct{})

	addLocale(config.JpMasterdataPath, "ja", config.MasterVersionJp, config.JpStartupKey)
	addLocale(config.GlMasterdataPath, "en", config.MasterVersionGl, config.GlStartupKey)
	addLocale(config.GlMasterdataPath, "zh", config.MasterVersionGl, config.GlStartupKey)
	addLocale(config.GlMasterdataPath, "ko", config.MasterVersionGl, config.GlStartupKey)

	for _, locale := range Locales {
		go locale.populate(syncChannel)
	}
	for i := len(Locales); i > 0; i-- {
		<-syncChannel
	}
	for _, local := range Locales {
		local.loadAsset()
	}

	finish := time.Now()
	log.Println("Finished loading databases in: ", finish.Sub(start))
	for language, locale := range Locales {
		gamedata.GamedataByLocale[language] = locale.Gamedata
		// because the order of has map is random, this instance is guaranteed to not
		// be a specific version, so don't depend on it
		gamedata.Instance = locale.Gamedata
	}
}
