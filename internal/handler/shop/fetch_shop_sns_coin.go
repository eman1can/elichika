package shop

import (
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"

	"github.com/gin-gonic/gin"
)

func fetchShopSnsCoin(ctx *gin.Context) {
	// there is no request body
	common.JsonResponse(ctx, &response.FetchShopSnsCoinResponse{})
}

func init() {
	server.AddHandler("/", "POST", "/shop/fetchShopSnsCoin", fetchShopSnsCoin)
}
