package bootstrap

import (
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"

	"github.com/gin-gonic/gin"
)

func getClearedPlatformAchievement(ctx *gin.Context) {
	common.JsonResponse(ctx, &response.GetClearedPlatformAchievementResponse{})
}

func init() {
	server.AddHandler("/", "POST", "/bootstrap/getClearedPlatformAchievement", getClearedPlatformAchievement)
}
