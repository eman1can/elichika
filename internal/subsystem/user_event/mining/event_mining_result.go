package mining

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
func resultEventScheduledHandler(userdata_db *xorm.Session, task server.ScheduledTask) {
	active := gamedata.Instance.EventActive
	eventIdInt, _ := strconv.Atoi(task.Params)
	if (active == nil) || (active.EventId != int32(eventIdInt)) ||
		(active.EventType != enum.EventTypeMining) || (active.ResultAt != task.Time) {
		log.Println("Warning: Failed to result event: ", task)
		return
	}
	user_info_trigger.CleanUpTriggerBasicByType(userdata_db, enum.InfoTriggerTypeEventMiningShowResult)
	eventMining := gamedata.Instance.EventMining[active.EventId]
	eventName := fmt.Sprintf("m.event_mining_title_%d", active.EventId)

	timePoint := time.Unix(task.Time, 0)
	// point ranking is awarded first, then voltage ranking
	// we will iterate over the point ranking, and get the voltage ranking that way
	pointRank := int32(0)
	pointRankingResults := GetPointRanking(userdata_db, active.EventId).GetRange(1, 1<<31-1)
	voltageRanking := GetVoltageRanking(userdata_db, active.EventId)
	for i, result := range pointRankingResults {
		if (i == 0) || (result.Score != pointRankingResults[i-1].Score) {
			pointRank = int32(i + 1)
		}
		session := userdata.GetBasicSession(userdata_db, timePoint, result.Id)
		rewardGroupId := eventMining.GetPointRankingReward(pointRank)
		for _, content := range gamedata.Instance.EventMiningReward[rewardGroupId] {
			user_present.AddPresent(session, client.PresentItem{
				Content:          *content,
				PresentRouteType: enum.PresentRouteTypeEventMiningPointRankingReward,
				PresentRouteId:   generic.NewNullable(active.EventId),
				ParamClient:      generic.NewNullable(strconv.Itoa(int(pointRank))),
				ParamServer: generic.NewNullable(client.LocalizedText{
					DotUnderText: eventName,
				}),
			})
		}
		voltageRank, isVoltageRanked := voltageRanking.TiedRankOf(result.Id)
		if isVoltageRanked {
			rewardGroupId := eventMining.GetVoltageRankingReward(int32(voltageRank))
			for _, content := range gamedata.Instance.EventMiningReward[rewardGroupId] {
				user_present.AddPresent(session, client.PresentItem{
					Content:          *content,
					PresentRouteType: enum.PresentRouteTypeEventMiningVoltageRankingReward,
					PresentRouteId:   generic.NewNullable(active.EventId),
					ParamClient:      generic.NewNullable(strconv.Itoa(int(pointRank))),
					ParamServer: generic.NewNullable(client.LocalizedText{
						DotUnderText: eventName, // this is resolved everytime user fetch present
					}),
				})
			}
		}
		user_info_trigger.AddTriggerBasic(session, client.UserInfoTriggerBasic{
			InfoTriggerType: enum.InfoTriggerTypeEventMiningShowResult,
			ParamInt:        generic.NewNullable(active.EventId),
		})
		session.Finalize()
	}

	// schedule the event actual end
	server.AddScheduledTask(server.ScheduledTask{
		Time:     active.EndAt,
		TaskName: "event_mining_end",
		Params:   task.Params,
	})

}

func init() {
	server.AddScheduledTaskHandler("event_mining_result", resultEventScheduledHandler)
}
