package sif_2_data_link

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func dataLink(ctx *gin.Context) {
	req := request.Sif2DataLinkRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	if req.IsPermission {
		common.JsonResponse(ctx, &response.Sif2DataLinkResponse{
			PassWord: "Kashikoi Kawaii Elichika",
		})
	} else {
		common.JsonResponse(ctx, &response.EmptyResponse{})
	}
}

func init() {
	server.AddHandler("/", "POST", "/sif2DataLink/dataLink", dataLink)
}
