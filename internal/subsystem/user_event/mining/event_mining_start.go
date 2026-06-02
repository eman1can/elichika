package mining

import (
	"fmt"
	"log"
	"strconv"

	"elichika/internal/enum"
	"elichika/internal/gamedata"
	"elichika/internal/server"
	"elichika/internal/utils"

	"xorm.io/xorm"
)

func ResetEventMiningResults(userdata_db *xorm.Session, eventId int32) {
	// Start the event.
	// This is only done once per event.
	// Because the event can be reused, this involve clearing out all the old record and trigger and stuff
	// The story progress will be kept
	_, err := userdata_db.Exec(fmt.Sprintf("UPDATE u_event_mining SET event_point = 0 WHERE event_master_id = %d", eventId))
	utils.CheckErr(err)
	_, err = userdata_db.Exec(fmt.Sprintf("DELETE FROM u_event_status WHERE event_id = %d", eventId))
	utils.CheckErr(err)
	_, err = userdata_db.Exec(fmt.Sprintf("DELETE FROM u_info_trigger_basic WHERE info_trigger_type = %d AND param_int = %d",
		enum.InfoTriggerTypeEventMiningFirstRuleDescription, eventId))
	utils.CheckErr(err)
	_, err = userdata_db.Exec(fmt.Sprintf("DELETE FROM u_info_trigger_basic WHERE info_trigger_type = %d AND param_int = %d",
		enum.InfoTriggerTypeEventMiningShowResult, eventId))
	utils.CheckErr(err)
	_, err = userdata_db.Exec("DELETE FROM u_event_mining_trade_product")
	utils.CheckErr(err)
	_, err = userdata_db.Exec("DELETE FROM u_event_mining_music_best_score")
	utils.CheckErr(err)
}

func StartScheduledEvent(eventId int32) *gamedata.EventActive {
	for _, schedule := range gamedata.Instance.EventSchedule {
		if schedule.EventId != eventId {
			continue
		}

		gamedata.Instance.EventActive = &gamedata.EventActive{
			EventId:   schedule.EventId,
			EventType: gamedata.Instance.Event[schedule.EventId].EventType,
			StartAt:   schedule.StartAt,
			ExpiredAt: schedule.ExpiredAt,
			ResultAt:  schedule.ResultAt,
			EndAt:     schedule.EndAt,
		}
		return gamedata.Instance.EventActive
	}

	return nil
}

func startEventScheduledHandler(userdata_db *xorm.Session, task server.ScheduledTask) {
	active := gamedata.Instance.EventActive
	eventIdInt, _ := strconv.Atoi(task.Params)
	eventId := int32(eventIdInt)

	if active != nil {
		log.Printf("Warning: Failed to start event: Event %d is already active!", active.EventId)
		return
	}

	event := gamedata.Instance.Event[eventId]
	if event.EventType != enum.EventTypeMining {
		log.Printf("Warning: Failed to start event: Event %d is not a mining event!", event.EventId)
		return
	}

	ResetEventMiningResults(userdata_db, eventId)

	active = StartScheduledEvent(eventId)
	if active == nil {
		log.Printf("Warning: Failed to start event: Event %d is not scheduled to start!", eventId)
		return
	}

	server.AddScheduledTask(server.ScheduledTask{
		Time:     active.ResultAt,
		TaskName: "event_mining_result",
		Params:   task.Params,
	})
}

func init() {
	server.AddScheduledTaskHandler("event_mining_start", startEventScheduledHandler)
}
