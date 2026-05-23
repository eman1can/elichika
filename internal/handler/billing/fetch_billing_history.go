package billing

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

// TODO(billing history): always return empty for now
// technically we can track usage but let's save that for later
func fetchBillingHistory(ctx *gin.Context) {
	req := request.BillingHistoryRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	common.JsonResponse(ctx, &response.BillingHistoryResponse{})
}

func init() {
	server.AddHandler("/", "POST", "/billing/fetchBillingHistory", fetchBillingHistory)
}
