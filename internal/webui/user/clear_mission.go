package user

import (
	"net/http"

	"elichika/internal/server"
	"elichika/internal/subsystem/user_mission"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

func clearMissions(ctx *gin.Context) {
	session := ctx.MustGet("session").(*userdata.Session)

	user_mission.FetchMission(session)

	for {
		var ids []int32
		for _, mission := range session.UserModel.UserMissionByMissionId.Map {
			if mission.IsReceivedReward {
				continue
			}
			mission.IsCleared = true
			mission.MissionCount = session.Gamedata.Mission[mission.MissionMId].MissionClearConditionCount
			ids = append(ids, mission.MissionMId)
		}
		if len(ids) == 0 {
			break
		}
		user_mission.ReceiveReward(session, ids)
	}
	session.Finalize()
	ctx.JSON(http.StatusOK, gin.H{})
}

func init() {
	server.AddHandler("/webui/user", "POST", "/clear_missions", clearMissions)
}
