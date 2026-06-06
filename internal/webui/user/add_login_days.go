package user

import (
	"net/http"

	"elichika/internal/server"

	"github.com/gin-gonic/gin"
)

func addLoginDays(ctx *gin.Context) {
	session, amount := getAddItemSession(ctx)

	if session != nil {
		session.UserStatus.LoginDays += min(amount, 10_000)

		session.Finalize()
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

func init() {
	server.AddHandler("/webui/user", "POST", "/add_login_days", addLoginDays)
}
