package shop

import (
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"

	"github.com/gin-gonic/gin"
)

func fetchShopPack(ctx *gin.Context) {
	common.JsonResponse(ctx, &response.FetchShopPackResponse{})
}

func init() {
	server.AddHandler("/", "POST", "/shop/fetchShopPack", fetchShopPack)
}
