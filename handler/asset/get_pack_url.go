package asset

import (
	"elichika/assetdata"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/config"
	"elichika/handler/common"
	"elichika/router"
	"elichika/utils"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func getPackUrl(ctx *gin.Context) {
	req := request.GetPackUrlRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	mv, exists := ctx.GetQuery("mv")
	utils.MustExist(exists)

	host := *config.Conf.ServerCdn

	resp := response.GetPackUrlResponse{}
	for _, pack := range req.PackNames.Slice {
		downloadData := assetdata.GetDownloadData(pack)
		if downloadData.IsEntireFile {
			if host == "elichika" {
				resp.UrlList.Append(fmt.Sprintf("%s/static/%s/%s/%s", host, mv, downloadData.Package, downloadData.File))
			} else {
				resp.UrlList.Append(fmt.Sprintf("%s/static/%s", host, downloadData.File))
			}
			continue
		}

		if host == "elichika" {
			resp.UrlList.Append(fmt.Sprintf("%s/static/%s/%s/%s&start=%d&size=%d", host, mv, downloadData.Package, downloadData.File, downloadData.Start, downloadData.Size))
		} else {
			resp.UrlList.Append(fmt.Sprintf("%s/%s", host, pack))
		}
	}

	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/", "POST", "/asset/getPackUrl", getPackUrl)
}
