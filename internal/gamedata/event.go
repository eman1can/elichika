package gamedata

import (
	"elichika/internal/utils"

	"xorm.io/xorm"
)

type Event struct {
	EventId      int32 `xorm:"id"`
	EventType    int32 `xorm:"event_type"`
	ReleaseOrder int32 `xorm:"release_order"`
	Available    bool  `xorm:"available"`
}

func loadEvent(gamedata *Gamedata) {
	gamedata.Event = make(map[int32]*Event)
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_event").Find(&gamedata.Event)
	})
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadEvent)
}
