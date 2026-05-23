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

func gradeUpCard(ctx *gin.Context) {
	req := request.GradeUpCardRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	user_training_tree.GradeUpCard(session, req.CardMasterId, req.ContentId)

	common.JsonResponse(ctx, response.GradeUpCardResponse{
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	server.AddHandler("/", "POST", "/trainingTree/gradeUpCard", gradeUpCard)
}
