package training_tree

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_training_tree"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func activateTrainingTreeCell(ctx *gin.Context) {
	req := request.ActivateTrainingTreeCellRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	user_training_tree.ActivateTrainingTreeCells(session, req.CardMasterId, req.CellMasterIds.Slice)

	session.Finalize()
	common.JsonResponse(ctx, &response.ActivateTrainingTreeCellResponse{
		UserCardTrainingTreeCellList: user_training_tree.GetUserTrainingTree(session, req.CardMasterId),
		UserModelDiff:                &session.UserModel,
	})
}

func init() {
	server.AddHandler("/", "POST", "/trainingTree/activateTrainingTreeCell", activateTrainingTreeCell)
}
