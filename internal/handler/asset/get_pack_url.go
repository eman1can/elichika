package asset

import (
	"encoding/json"
	"fmt"

	"elichika/internal/assetdata"
	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/config"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func getPackUrl(ctx *gin.Context) {
	req := request.GetPackUrlRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	host := *config.Conf.ServerCdn

	resp := response.GetPackUrlResponse{}
	for _, pack := range req.PackNames.Slice {
		downloadData := assetdata.GetDownloadData(pack)
		if downloadData.IsEntireFile {
			resp.UrlList.Append(fmt.Sprintf("%s/static/%s", host, downloadData.File))
			continue
		}

		if host == config.DefaultServerCdn {
			resp.UrlList.Append(fmt.Sprintf("%s/static/%s?start=%d&size=%d", host, downloadData.File, downloadData.Start, downloadData.Size))
		} else {
			resp.UrlList.Append(fmt.Sprintf("%s/%s", host, pack))
		}
	}

	common.JsonResponse(ctx, &resp)
}

func init() {
	server.AddHandler("/", "POST", "/asset/getPackUrl", getPackUrl)
}
