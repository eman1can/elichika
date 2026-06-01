package user

import (
	"net/http"

	"elichika/internal/item"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_content"

	"github.com/gin-gonic/gin"
)

func addSnsCoin(ctx *gin.Context) {
	session, amount := getAddItemSession(ctx)

	if session != nil {
		user_content.AddContent(session, item.StarGem.Amount(min(amount, 1_000_000)))

		session.Finalize()
		ctx.Status(http.StatusOK)
	}
}

func init() {
	server.AddHandler("/webui/user", "POST", "/add_sns_coin", addSnsCoin)
}
