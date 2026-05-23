package live_partners

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_social"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func setLivePartner(ctx *gin.Context) {
	req := request.SetLivePartnerCardRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	user_social.SetLivePartnerCard(session, req.LivePartnerCategoryId, req.CardMasterId)

	common.JsonResponse(ctx, response.EmptyResponse{})
}

func init() {
	server.AddHandler("/", "POST", "/livePartners/setLivePartner", setLivePartner)
}
