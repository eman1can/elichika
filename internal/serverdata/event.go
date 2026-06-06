package serverdata

import (
	"fmt"
	"log"
	"os"

	"elichika/internal/config"
	"elichika/internal/enum"
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

	Title       string `xorm:"title" json:"title"`
	Description string `xorm:"description" json:"description"`

	BannerNoticeSmallAssetPath *string `xorm:"banner_notice_s" json:"banner_notice_s"`
	BannerNoticeLargeAssetPath *string `xorm:"banner_notice_l" json:"banner_notice_l"`
}

func (event *Event) populate(session *xorm.Session) {
	switch event.EventType {
	case enum.EventTypeMining:
		event.Title = fmt.Sprintf("m.event_mining_title_%d", event.EventId)
		event.Description = "Mining Event"
	case enum.EventTypeMarathon:
		event.Title = fmt.Sprintf("m.event_marathon_title_%d", event.EventId)
		event.Description = "Marathon Event"
	default:
		event.Title = fmt.Sprintf("Unknown Event: %d", event.EventId)
		event.Description = "Unknown Event"
	}
}

func initEvent(session *xorm.Session) {
	path := config.AssetPath + "event/events.json"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Panic("Could not find event JSON file.")
		return
	}

	var events []*Event
	parser.ParseJson(path, &events)

	for _, event := range events {
		event.populate(session)
	}

	_, err := session.Table("s_event").Insert(events)
	utils.CheckErr(err)

	log.Printf("Loaded %d events", len(events))
}

func init() {
	addTable("s_event", Event{}, initEvent)
}
