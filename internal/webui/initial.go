package webui

import (
	"elichika/internal/locale"
	"elichika/internal/server"
	_ "elichika/internal/webui/agnostic"
	_ "elichika/internal/webui/auth"
	_ "elichika/internal/webui/user"

	"github.com/gin-gonic/gin"
)

func webuiInitial(ctx *gin.Context) {
	if server.IsShutdown() {
		ctx.Abort()
		return
	}
	server.StartConnection()
	defer server.FinishConnection()

	lang, _ := ctx.GetQuery("l")
	if lang == "" {
		lang = "en"
	}
	ctx.Set("locale", locale.Locales[lang])
	ctx.Set("gamedata", locale.Locales[lang].Gamedata)
	ctx.Set("dictionary", locale.Locales[lang].Dictionary)

	ctx.Next()
}

func init() {
	server.AddInitialHandler("/webui", webuiInitial)
}
