package event

import (
	"strings"

	"elichika/internal/server"
	"elichika/internal/serverstate"
)

var isDirectEventChanging bool
var targetedEventId int32

// select the event directly instead of waiting for the scheduler
func ChangeEvent(eventId int32) {
	isDirectEventChanging = true
	targetedEventId = eventId

	server.ForceRun(nil, func(task serverstate.ScheduledTask) (bool, bool) {
		if task.TaskName == "event_marathon_start" {
			return true, true // run the task and then stop
		} else if task.TaskName == "event_mining_start" {
			return true, true // run the task and then stop
		}
		return strings.HasPrefix(task.TaskName, "event"), false // run the task if it's event related, to clean up existing events, then stop
	})
	isDirectEventChanging = false
}
