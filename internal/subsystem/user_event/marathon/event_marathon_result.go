package marathon

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/gamedata"
	"elichika/internal/generic"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_info_trigger"
	"elichika/internal/subsystem/user_present"
	"elichika/internal/userdata"

	"xorm.io/xorm"
)

// finish the event and pay out the reward for everyone who participated
func resultEventScheduledHandler(userdataDb *xorm.Session, task server.ScheduledTask) {
	active := gamedata.Instance.EventActive

	eventIdInt, _ := strconv.Atoi(task.Params)
	eventId := int32(eventIdInt)
	if (active == nil) || (active.EventId != eventId) || (active.EventType != enum.EventTypeMarathon) || (active.ResultAt < task.Time) {
		log.Println("Warning: Failed to result event: ", task)
		return
	}

	results := GetRanking(userdataDb, eventId).GetRange(1, 1<<31-1)
	eventMarathon := gamedata.Instance.EventMarathon[eventId]

	rank := int32(0)
	timePoint := time.Unix(task.Time, 0)
	user_info_trigger.CleanUpTriggerBasicByType(userdataDb, enum.InfoTriggerTypeEventMarathonShowResult)
	for i, result := range results {
		if (i == 0) || (result.Score != results[i-1].Score) {
			rank = int32(i + 1)
		}
		session := userdata.GetBasicSession(userdataDb, timePoint, result.Id)
		rewardGroupId := eventMarathon.GetRankingReward(rank)
		for _, content := range gamedata.Instance.EventMarathonReward[rewardGroupId] {
			user_present.AddPresent(session, client.PresentItem{
				Content:          *content,
				PresentRouteType: enum.PresentRouteTypeEventMarathonRankingReward,
				PresentRouteId:   generic.NewNullable(eventMarathon.EventId),
				ParamClient:      generic.NewNullable(strconv.Itoa(int(rank))),
				ParamServer: generic.NewNullable(client.LocalizedText{
					DotUnderText: fmt.Sprintf("m.event_marathon_title_%d", eventId),
				}),
			})
		}

		user_info_trigger.AddTriggerBasic(session, client.UserInfoTriggerBasic{
			InfoTriggerType: enum.InfoTriggerTypeEventMarathonShowResult,
			ParamInt:        generic.NewNullable(eventId),
		})
		session.Finalize()
	}

	// schedule the event actual end
	server.AddScheduledTask(server.ScheduledTask{
		Time:     active.EndAt,
		TaskName: "event_marathon_end",
		Params:   task.Params,
	})
}

func init() {
	server.AddScheduledTaskHandler("event_marathon_result", resultEventScheduledHandler)
}
