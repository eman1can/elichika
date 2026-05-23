package asset

import (
	"net/http"
	"os"

	"elichika/internal/server"
	"elichika/internal/utils"

	"strconv"

	"github.com/gin-gonic/gin"
)

func staticApi(ctx *gin.Context) {
	file := ctx.Param("path")
	path := "static" + file

	info, err := os.Stat(path)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	startString, startStringExist := ctx.GetQuery("start")
	sizeString, sizeStringExist := ctx.GetQuery("size")

	if startStringExist && sizeStringExist {
		start, err := strconv.Atoi(startString)
		utils.CheckErr(err)
		size, err := strconv.Atoi(sizeString)
		utils.CheckErr(err)

		sendRange(ctx, path, start, size)
	} else {
		sendFile(ctx, path, info.Size())
	}
}

func init() {
	server.AddHandler("/static", "GET", "/*path", staticApi)
}
