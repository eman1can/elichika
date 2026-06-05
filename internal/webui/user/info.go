package user

import (
	"net/http"

	"elichika/internal/server"
	"elichika/internal/subsystem/user_status"
	"elichika/internal/subsystem/user_story_main"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

type WebUIUserInfoResponse struct {
	GameMoney        int32 `json:"money"`
	Experience       int32 `json:"experience"`
	LivePoints       int32 `json:"lp"`
	ActivityPoints   int32 `json:"ap"`
	SnsCoin          int32 `json:"sns_coin"`
	SubscriptionCoin int32 `json:"subscription_coin"`
	LoginDays        int32 `json:"login_days"`
	StoryFinished    bool  `json:"story_finished"`
}

func getUserInfo(ctx *gin.Context) {
	var resp WebUIUserInfoResponse

	session := ctx.MustGet("session").(*userdata.Session)

	resp.GameMoney = session.UserStatus.GameMoney
	resp.Experience = session.UserStatus.Exp
	resp.ActivityPoints = session.UserStatus.ActivityPointCount
	resp.LivePoints = user_status.GetUserLivePoints(session)
	resp.SnsCoin = session.UserStatus.FreeSnsCoin
	resp.SubscriptionCoin = session.UserStatus.SubscriptionCoin
	resp.LoginDays = session.UserStatus.LoginDays
	resp.StoryFinished = user_story_main.AllStoryFinished(session)

	ctx.JSON(http.StatusOK, resp)
}

func init() {
	server.AddHandler("/webui/user", "GET", "/info", getUserInfo)
}
