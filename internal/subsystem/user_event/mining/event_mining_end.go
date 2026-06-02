package mining

import (
	"log"
	"strconv"

	"elichika/internal/enum"
	"elichika/internal/gamedata"
	"elichika/internal/server"

	"xorm.io/xorm"
)

// finish the event and pay out the reward for everyone who participated
func endEventScheduledHandler(userdata_db *xorm.Session, task server.ScheduledTask) {
	active := gamedata.Instance.EventActive
	eventIdInt, _ := strconv.Atoi(task.Params)
	eventId := int32(eventIdInt)
	if (active == nil) || (active.EventId != eventId) || (active.EventType != enum.EventTypeMining) || (active.EndAt != task.Time) {
		log.Println("Warning: Failed to end event: ", task)
		return
	}
	// no actual clean up is necessary, we just need to remove the ranking object
	ResetVoltageRanking()
	ResetPointRanking()

	// TODO(event): Add config for other options once we have more than 1 event
	server.AddScheduledTask(server.ScheduledTask{
		Time:     active.EndAt + 1,
		TaskName: "event_auto_scheduler",
	})
}

func init() {
	server.AddScheduledTaskHandler("event_mining_end", endEventScheduledHandler)
}
