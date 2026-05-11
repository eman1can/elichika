package asset

import (
	"elichika/log"
	"elichika/utils"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func sendFile(ctx *gin.Context, fileName string, size int64) {
	ctx.Header("Content-Length", fmt.Sprint(size))
	ctx.Header("Content-Type", "application/octet-stream")

	ctx.File(fileName)
}

func sendRange(ctx *gin.Context, fileName string, start, size int) {
	ctx.Header("Content-Length", fmt.Sprint(size))
	ctx.Header("Content-Type", "application/octet-stream")

	buffer := make([]byte, 1024)
	f, err := os.Open(fileName)
	utils.CheckErr(err)
	defer f.Close()
	_, err = f.Seek(int64(start), io.SeekStart)
	utils.CheckErr(err)
	for ; size > 0; size -= 1024 {
		count, err := f.Read(buffer)
		utils.CheckErr(err)
		if count > size {
			count = size
		} else if (count < 1024) && (count < size) {
			log.Panic("wrong requested range")
		}
		ctx.Writer.Write(buffer[:count])
	}
}
