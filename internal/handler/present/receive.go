package present

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_present"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func receive(ctx *gin.Context) {
	req := request.ReceivePresentRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	resp := response.ReceivePresentResponse{
		UserModelDiff: &session.UserModel,
	}

	for _, id := range req.Ids.Slice {
		user_present.Receive(session, id, &resp)
	}

	session.Finalize()
	resp.PresentItems = user_present.FetchPresentItems(session)
	resp.PresentHistoryItems = user_present.FetchPresentHistoryItems(session)
	resp.PresentCount = user_present.FetchPresentCount(session)
	common.JsonResponse(ctx, &resp)
}

func init() {
	server.AddHandler("/", "POST", "/present/receive", receive)
}
