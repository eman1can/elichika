package webui

import (
	"elichika/internal/locale"
	"elichika/internal/server"
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

	platform, _ := ctx.GetQuery("p")
	if platform == "" {
		platform = "a"
	}

	if platform == "a" {
		ctx.Set("assetdata", locale.Locales[lang].AssetdataAndroid)
	} else {
		ctx.Set("assetdata", locale.Locales[lang].AssetdataIos)
	}

	ctx.Next()
}

func init() {
	server.AddInitialHandler("/webui", webuiInitial)
	server.AddInitialHandler("/webui/user", webuiInitial)
	server.AddInitialHandler("/webui/admin", webuiInitial)
}
