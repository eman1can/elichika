package handler

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_rule_description"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func saveRuleDescription(ctx *gin.Context) {
	req := request.SaveRuleDescriptionRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	// response with user model
	session := ctx.MustGet("session").(*userdata.Session)

	for _, ruleDescriptionId := range req.RuleDescriptionMasterIds.Slice {
		user_rule_description.UpdateUserRuleDescription(session, ruleDescriptionId)
	}

	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	server.AddHandler("/", "POST", "/ruleDescription/saveRuleDescription", saveRuleDescription)
}
