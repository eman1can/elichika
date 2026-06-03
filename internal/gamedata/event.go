package gamedata

import (
	"elichika/internal/serverdata"
	"elichika/internal/utils"

	"xorm.io/xorm"
)

func loadEvent(gamedata *Gamedata) {
	gamedata.Event = make(map[int32]*serverdata.Event)

	var err error
	gamedata.ServerdataDb.Do(func(session *xorm.Session) {
		err = session.Table("s_event").Find(&gamedata.Event)
	})
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadEvent)
}
