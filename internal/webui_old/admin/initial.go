package admin

import (
	"log"

	"elichika/internal/server"
	"elichika/internal/utils"

	"bytes"
	"encoding/base64"
	"strings"

	"github.com/gin-gonic/gin"
)

func adminInitial(ctx *gin.Context) {
	if ctx.Request.Method == "POST" {
		form, err := ctx.MultipartForm()
		utils.CheckErr(err)
		ctx.Set("form", form)
		if !strings.HasPrefix(ctx.Request.URL.String(), "/webui/admin/login") {
			sessionKey, err := base64.StdEncoding.DecodeString(form.Value["admin_session_key"][0])
			utils.CheckErr(err)
			if !bytes.Equal(sessionKey, adminSessionKey) {
				log.Panic("wrong session key")
			}
		}
	}
	ctx.Next()
}

func init() {
	server.AddInitialHandler("/webui/admin", adminInitial)
}
