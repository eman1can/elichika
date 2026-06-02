package marathon

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

func StartEventMarathon(userdata_db *xorm.Session, eventId int32) {
	// Start the event.
	// This is only done once per event.
	// Because the event can be reused, this involve clearing out all the old record and trigger and stuff
	// The story progress will be kept
	_, err := userdata_db.Exec(fmt.Sprintf("UPDATE u_event_marathon SET event_point = 0 WHERE event_master_id = %d", eventId))
	utils.CheckErr(err)
	_, err = userdata_db.Exec(fmt.Sprintf("DELETE FROM u_event_status WHERE event_id = %d", eventId))
	utils.CheckErr(err)
	_, err = userdata_db.Exec(fmt.Sprintf("DELETE FROM u_info_trigger_basic WHERE info_trigger_type = %d AND param_int = %d",
		enum.InfoTriggerTypeEventMarathonFirstRuleDescription, eventId))
	utils.CheckErr(err)
	_, err = userdata_db.Exec(fmt.Sprintf("DELETE FROM u_info_trigger_basic WHERE info_trigger_type = %d AND param_int = %d",
		enum.InfoTriggerTypeEventMarathonShowResult, eventId))
	utils.CheckErr(err)
	_, err = userdata_db.Exec(fmt.Sprintf("DELETE FROM u_content WHERE content_type = %d AND content_id = %d",
		enum.ContentTypeEventMarathonBooster, gamedata.Instance.EventMarathon[eventId].BoosterItemId))
	utils.CheckErr(err)
}

func startEventScheduledHandler(userdata_db *xorm.Session, task server.ScheduledTask) {
	active := gamedata.Instance.EventActive
	eventIdInt, _ := strconv.Atoi(task.Params)

	if (active == nil) || (active.EventId != int32(eventIdInt)) || (active.EventType != enum.EventTypeMarathon) || (active.StartAt > task.Time) {

		log.Println("Warning: Failed to start event: ", task)
		return
	}
	// this will be scheduled by an event scheduler, and called when the event is ready to start
	StartEventMarathon(userdata_db, active.EventId)

	// schedule the event payout and stuff
	server.AddScheduledTask(server.ScheduledTask{
		Time:     active.ResultAt,
		TaskName: "event_marathon_result",
		Params:   task.Params,
	})
}

func init() {
	server.AddScheduledTaskHandler("event_marathon_start", startEventScheduledHandler)
}
