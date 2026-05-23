package tower

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_tower"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func recoveryTowerCardUsedAll(ctx *gin.Context) {
	req := request.RecoveryTowerCardUsedRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	resp := response.RecoveryTowerCardUsedResponse{
		TowerCardUsedCountRows: user_tower.GetUserTowerCardUsedList(session, req.TowerId),
		UserModelDiff:          &session.UserModel,
	}
	for i := range resp.TowerCardUsedCountRows.Slice {
		resp.TowerCardUsedCountRows.Slice[i].UsedCount = 0
		resp.TowerCardUsedCountRows.Slice[i].RecoveredCount = 0
		user_tower.UpdateUserTowerCardUsed(session, req.TowerId, resp.TowerCardUsedCountRows.Slice[i])
	}

	common.JsonResponse(ctx, &resp)
}

func init() {
	server.AddHandler("/", "POST", "/tower/recoveryTowerCardUsedAll", recoveryTowerCardUsedAll)
}
