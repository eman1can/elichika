package challenge

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_beginner_challenge"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func receiveRewardBeginner(ctx *gin.Context) {
	req := request.ChallengeBeginnerRewardRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	common.JsonResponse(ctx, user_beginner_challenge.ReceiveRewardBeginner(session, req.ChallengeId, req.ChallengeCellId))
}

func init() {
	server.AddHandler("/", "POST", "/challenge/receiveRewardBeginner", receiveRewardBeginner)
}
