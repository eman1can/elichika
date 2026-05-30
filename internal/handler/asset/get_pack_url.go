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

func getHostUrl(ctx *gin.Context) string {
	host := *config.Conf.ServerCdn

	if host == "elichika" {
		actualHost := ctx.Request.Host

		actualProto := "http"
		if ctx.Request.TLS != nil {
			actualProto = "https"
		}

		// if the connection is forwarded, we need to return the forwarded host instead
		forwardedHost, hostExists := ctx.Request.Header["X-Forwarded-Host"]
		forwardedProto, protoExists := ctx.Request.Header["X-Forwarded-Proto"]
		if hostExists && len(forwardedHost) > 0 {
			actualHost = forwardedHost[0]
		}
		if protoExists && len(forwardedProto) > 0 {
			actualProto = forwardedProto[0]
		}

		return actualProto + "://" + actualHost
	}

	return host
}

func getPackUrl(ctx *gin.Context) {
	req := request.GetPackUrlRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	host := getHostUrl(ctx)
	isElichika := *config.Conf.ServerCdn == "elichika"

	resp := response.GetPackUrlResponse{}
	for _, pack := range req.PackNames.Slice {
		downloadData := assetdata.GetDownloadData(pack)
		if downloadData.IsEntireFile {
			resp.UrlList.Append(fmt.Sprintf("%s/static/%s", host, downloadData.File))
			continue
		}

		if isElichika {
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
