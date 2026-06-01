package user

import (
	"net/http"

	"elichika/internal/server"
	"elichika/internal/subsystem/user_status"

	"github.com/gin-gonic/gin"
)

func addActivityPoints(ctx *gin.Context) {
	session, amount := getAddItemSession(ctx)

	if session != nil {
		user_status.AddUserActivityPoints(session, amount)

		session.Finalize()
		ctx.Status(http.StatusOK)
	}
}

func init() {
	server.AddHandler("/webui/user", "POST", "/add_activity_point", addActivityPoints)
}
