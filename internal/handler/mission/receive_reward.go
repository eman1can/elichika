package mission

import (
	"log"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_mission"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func receiveReward(ctx *gin.Context) {
	req := request.MissionRewardRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	resp := user_mission.ReceiveReward(session, req.MissionIds.Slice)

	switch resp.(type) {
	case response.MissionReceiveResponse:
		common.JsonResponse(ctx, &resp)
	case response.MissionReceiveErrorResponse:
		common.AlternativeJsonResponse(ctx, &resp)
	default:
		log.Panic("not supported")
	}
}

func init() {
	server.AddHandler("/", "POST", "/mission/receiveReward", receiveReward)
}
