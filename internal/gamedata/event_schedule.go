package gamedata

import (
	"time"

	"elichika/internal/serverstate"
	"elichika/internal/utils"

	"xorm.io/xorm"
)

type EventActive struct {
	EventId   int32
	EventType int32
	StartAt   int64
	ExpiredAt int64
	ResultAt  int64
	EndAt     int64
}

func loadEventSchedule(gamedata *Gamedata) {
	var err error
	serverstate.Database.Do(func(session *xorm.Session) {
		err = session.Table("s_event_active").Find(&gamedata.EventSchedule)
	})
	utils.CheckErr(err)

	current := time.Now().Unix()
	gamedata.EventActive = nil
	for _, schedule := range gamedata.EventSchedule {
		if schedule.StartAt <= current && schedule.EndAt > current {
			gamedata.EventActive = &EventActive{
				EventId:   schedule.EventId,
				EventType: gamedata.Event[schedule.EventId].EventType,
				StartAt:   schedule.StartAt,
				ExpiredAt: schedule.ExpiredAt,
				ResultAt:  schedule.ResultAt,
				EndAt:     schedule.EndAt,
			}
		}
	}
}

func init() {
	addLoadFunc(loadEventSchedule)
	addPrequisite(loadEventSchedule, loadEvent)
}
