package data_link

import (
	"elichika/internal/client"
	"elichika/internal/handler/common"
	"elichika/internal/server"

	"github.com/gin-gonic/gin"
)

func fetchDataLinks(ctx *gin.Context) {
	// there's no request body
	common.JsonResponse(ctx, client.LinkedInfo{
		IsTakeOverIdLinked: true,
	})
}

func init() {
	server.AddHandler("/", "POST", "/dataLink/fetchDataLinks", fetchDataLinks)
}
