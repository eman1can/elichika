package user

import (
	"bytes"
	"encoding/base64"
	"log"
	"strconv"
	"strings"

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
	if ctx.Request.Method == "POST" {
		log.Println("Accepting: ", ctx.Request.URL.String())
		form, err := ctx.MultipartForm()
		utils.CheckErr(err)
		ctx.Set("form", form)
		userIdString, exist := form.Value["user_id"]
		if exist {
			userId, err := strconv.Atoi(userIdString[0])
			utils.CheckErr(err)
			ctx.Set("user_id", userId)
			session = userdata.GetSession(ctx, int32(userId))
			if !strings.HasPrefix(ctx.Request.URL.String(), "/webui/user/login") {
				sessionKey, err := base64.StdEncoding.DecodeString(form.Value["user_session_key"][0])
				utils.CheckErr(err)
				if !bytes.Equal(sessionKey, session.SessionKey()) {
					log.Panic("wrong session key")
				}
			} else {
				session.SessionType = userdata.SessionTypeLogin
			}
		}
		ctx.Set("session", session)
	}
	ctx.Next()
}

func init() {
	server.AddInitialHandler("/webui/user", userInitial)
}
