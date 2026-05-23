package user

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/config"
	"elichika/internal/handler/common"
	"elichika/internal/item"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_content"
	"elichika/internal/subsystem/user_status"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func recoverLp(ctx *gin.Context) {
	req := request.RecoverLPRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	// TODO(hardcode): technically this is defined in m_recovery_lp
	// and SnsCoinForLpRecover is in m_constant
	switch req.ContentId {
	case item.ShowCandy50.ContentId:
		user_status.AddUserLp(session, req.Count.Value*50)
		if config.Conf.ResourceConfig().ConsumeMiscItems {
			user_content.RemoveContent(session, item.ShowCandy50.Amount(req.Count.Value))
		}
	case item.ShowCandy100.ContentId:
		user_status.AddUserLp(session, req.Count.Value*100)
		if config.Conf.ResourceConfig().ConsumeMiscItems {
			user_content.RemoveContent(session, item.ShowCandy100.Amount(req.Count.Value))
		}
	case item.StarGem.ContentId:
		user_status.AddUserLp(session, req.Count.Value*100)
		if config.Conf.ResourceConfig().ConsumeMiscItems {
			user_content.RemoveContent(session, item.StarGem.Amount(req.Count.Value*10))
		}
	}

	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	server.AddHandler("/", "POST", "/user/recoverLp", recoverLp)
}
