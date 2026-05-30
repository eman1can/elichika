package asset

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func sendFile(ctx *gin.Context, path string, size int64) {
	ctx.Header("Content-Length", fmt.Sprint(size))
	ctx.Header("Content-Type", "application/octet-stream")

	ctx.Status(http.StatusOK)
	ctx.File(path)
}

func sendRange(ctx *gin.Context, path string, start, size int) {
	ctx.Header("Content-Length", fmt.Sprint(size))
	ctx.Header("Content-Type", "application/octet-stream")

	f, err := os.Open(path)
	utils.CheckErr(err)
	defer f.Close()
	_, err = f.Seek(int64(start), io.SeekStart)
	utils.CheckErr(err)

	buffer := make([]byte, 64*1024)
	remaining := size
	for remaining > 0 {
		toRead := remaining
		if toRead > len(buffer) {
			toRead = len(buffer)
		}
		count, err := f.Read(buffer[:toRead])
		utils.CheckErr(err)
		ctx.Writer.Write(buffer[:count])
		remaining -= count
	}
	ctx.Status(http.StatusPartialContent)
}
