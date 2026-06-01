package user

import (
	"bytes"
	"encoding/base64"
	"log"
	"strconv"

	"elichika/internal/locale"
	"elichika/internal/server"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func userInitial(ctx *gin.Context) {
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

	var session *userdata.Session
	defer func() { session.Close() }()

	userId, exist := ctx.GetQuery("u")
	if exist {
		userId, err := strconv.Atoi(userId)
		utils.CheckErr(err)

		ctx.Set("user_id", userId)
		session = userdata.GetSession(ctx, int32(userId))

		sessionHeader := ctx.GetHeader("X-Session-Key")
		if sessionHeader != "" {
			sessionKey, err := base64.StdEncoding.DecodeString(sessionHeader)
			utils.CheckErr(err)
			if !bytes.Equal(sessionKey, session.SessionKey()) {
				log.Panic("wrong session key")
			}
		}
	}
	ctx.Set("session", session)
	ctx.Next()
}

func init() {
	server.AddInitialHandler("/webui/user", userInitial)
}
