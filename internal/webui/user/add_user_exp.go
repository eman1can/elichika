package user

import (
	"net/http"

	"elichika/internal/server"
	"elichika/internal/subsystem/user_status"

	"github.com/gin-gonic/gin"
)

func addExperience(ctx *gin.Context) {
	session, amount := getAddItemSession(ctx)

	if session != nil {
		user_status.AddUserExp(session, amount)

		session.Finalize()
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

func init() {
	server.AddHandler("/webui/user", "POST", "/add_user_experience", addExperience)
}
