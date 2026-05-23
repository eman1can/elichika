package still

import (
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"

	"github.com/gin-gonic/gin"
)

func fetch(ctx *gin.Context) {
	// there is no request body

	common.JsonResponse(ctx, &response.FetchStillResponse{})
}

func init() {
	server.AddHandler("/", "POST", "/still/fetch", fetch)
}
