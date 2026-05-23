package event_mining

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/event"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

// response: LikeEventMiningPanelResponse
func likeEventMiningPanel(ctx *gin.Context) {
	req := request.LikeEventMiningPanelRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)
	common.JsonResponse(ctx, event.LikeEventMiningPanel(session, req.EventId, req.ThumbnailCellId))
}

func init() {
	server.AddHandler("/", "POST", "/eventMining/likeEventMiningPanel", likeEventMiningPanel)
}
