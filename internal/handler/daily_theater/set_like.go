package daily_theater

import (
	"encoding/json"

	"elichika/internal/client"
	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func setLike(ctx *gin.Context) {
	req := request.DailyTheaterSetLikeRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	session.UserModel.UserDailyTheaterByDailyTheaterId.Set(
		req.DailyTheaterId,
		client.UserDailyTheater{
			DailyTheaterId: req.DailyTheaterId,
			IsLiked:        req.IsLike,
		})

	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	server.AddHandler("/", "POST", "/dailyTheater/setLike", setLike)
}
