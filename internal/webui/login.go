package webui

import (
	"elichika/internal/config"
	"elichika/internal/enum"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_authentication"
	"elichika/internal/userdata"
	"elichika/internal/utils"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"

	"crypto/rand"
	"crypto/subtle"

	"github.com/gin-gonic/gin"
)

var sessionKey []byte

func randomKey() []byte {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	utils.CheckErr(err)
	return b
}

func newSessionKey() {
	sessionKey = randomKey()
}

func login(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	utils.CheckErr(err)

	username := form.Value["username"][0]
	password := form.Value["password"][0]

	var s string
	resp := Response{}

	if username == "admin" {
		if subtle.ConstantTimeCompare([]byte(*config.Conf.AdminPassword), []byte(password)) == 1 {
			newSessionKey()
			s = base64.StdEncoding.EncodeToString(sessionKey)
			resp.Response = &s
		} else {
			s = "Wrong password!"
			resp.Error = &s
		}
	} else {
		userId, parseErr := strconv.Atoi(username)
		if parseErr != nil {
			s = "Invalid user ID!"
			resp.Error = &s
		} else {
			session := userdata.GetSession(ctx, int32(userId))
			defer session.Close()
			if session == nil {
				s = "User doesn't exist!"
				resp.Error = &s
			} else {
				session.SessionType = userdata.SessionTypeLogin
				if !user_authentication.CheckPassWord(session, password) {
					s = "Wrong password!"
					resp.Error = &s
				} else if session.UserStatus.TutorialPhase != enum.TutorialPhaseTutorialEnd {
					s = "Finish the tutorial (in game) first before using the WebUI!"
					resp.Error = &s
				} else {
					s = base64.StdEncoding.EncodeToString(session.SessionKey())
					resp.Response = &s
				}
			}
		}
	}

	jsonBytes, err := json.Marshal(resp)
	utils.CheckErr(err)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, string(jsonBytes))
}

func init() {
	server.AddHandler("/webui", "POST", "/login", login)
}
