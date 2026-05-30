package asset

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"elichika/internal/config"
	"elichika/internal/server"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func downloadFromProxy(path string, pack string) error {
	resp, err := http.Get(fmt.Sprintf("%s/static/%s", config.DefaultProxyCdn, pack))
	if err != nil {
		return err
	}

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

type StaticFileRequest struct {
	Start *int `form:"start"`
	Size  *int `form:"size"`
}

func staticApi(ctx *gin.Context) {
	path := filepath.Join(config.StaticDataPath, "packs", ctx.Param("path"))

	info, err := os.Stat(path)
	if err != nil {
		path = filepath.Join(config.StaticDataPath, ctx.Param("path"))
		info, err = os.Stat(path)
	}

	if err != nil {
		log.Println("Download", path, "from proxy", *config.Conf.StaticProxyCdn)
		err = downloadFromProxy(path, ctx.Param("path"))
	}

	if err == nil {
		req := StaticFileRequest{}
		err := ctx.ShouldBindQuery(&req)
		utils.CheckErr(err)

		if req.Start != nil && req.Size != nil {
			sendRange(ctx, path, *req.Start, *req.Size)
		} else {
			sendFile(ctx, path, info.Size())
		}
	}

	ctx.Status(http.StatusNotFound)
}

func init() {
	server.AddHandler("/static", "GET", "/*path", staticApi)
}
