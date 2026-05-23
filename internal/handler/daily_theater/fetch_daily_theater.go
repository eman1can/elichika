package daily_theater

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_daily_theater"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func fetchDailyTheater(ctx *gin.Context) {
	req := request.FetchDailyTheaterRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	successResponse, failureRespose := user_daily_theater.FetchDailyTheater(session, req.DailyTheaterId)
	if successResponse == nil {
		common.AlternativeJsonResponse(ctx, failureRespose)
	} else {
		common.JsonResponse(ctx, successResponse)
	}
}

func init() {
	server.AddHandler("/", "POST", "/dailyTheater/fetchDailyTheater", fetchDailyTheater)
}
