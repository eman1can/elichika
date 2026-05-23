package accessory

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_accessory"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func powerUp(ctx *gin.Context) {
	req := request.AccessoryPowerUpRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	resp := user_accessory.LevelUpAccessory(session, req.UserAccessoryId, req.PowerUpAccessoryIds, req.AccessoryLevelUpItems)

	common.JsonResponse(ctx, &resp)
}

func init() {
	server.AddHandler("/", "POST", "/accessory/powerUp", powerUp)
}
