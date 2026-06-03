package serverdata

import (
	"log"
	"os"

	"elichika/internal/config"
	"elichika/internal/parser"
	"elichika/internal/utils"

	"xorm.io/xorm"
)

type Event struct {
	EventId       int32 `xorm:"pk 'id'" json:"id"`
	EventType     int32 `xorm:"event_type" json:"event_type"`
	ReleaseOrder  int32 `xorm:"release_order" json:"release_order"`
	Available     bool  `xorm:"available" json:"available"`
	GachaMasterId int32 `xorm:"gacha_master_id" json:"gacha_master_id"`

	BannerNoticeSmallAssetPath *string `xorm:"banner_notice_s" json:"banner_notice_s"`
	BannerNoticeLargeAssetPath *string `xorm:"banner_notice_l" json:"banner_notice_l"`
}

func initEvent(session *xorm.Session) {
	path := config.AssetPath + "event/events.json"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Panic("Could not find event JSON file.")
		return
	}

	var events []Event
	parser.ParseJson(path, &events)
	_, err := session.Table("s_event").Insert(events)
	utils.CheckErr(err)

	log.Printf("Loaded %d events", len(events))
}

func init() {
	addTable("s_event", Event{}, initEvent)
}
