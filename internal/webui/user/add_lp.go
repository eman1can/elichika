package user

import (
	"net/http"

	"elichika/internal/server"
	"elichika/internal/subsystem/user_status"

	"github.com/gin-gonic/gin"
)

func addLivePoints(ctx *gin.Context) {
	session, amount := getAddItemSession(ctx)

	if session != nil {
		user_status.AddUserLivePoints(session, min(amount, 10_000))

		session.Finalize()
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

func init() {
	server.AddHandler("/webui/user", "POST", "/add_live_point", addLivePoints)
}
