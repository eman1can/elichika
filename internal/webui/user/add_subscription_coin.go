package user

import (
	"net/http"

	"elichika/internal/item"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_content"

	"github.com/gin-gonic/gin"
)

func addSubscriptionCoin(ctx *gin.Context) {
	session, amount := getAddItemSession(ctx)

	if session != nil {
		user_content.AddContent(session, item.MemberCoin.Amount(min(amount, 100_000)))

		session.Finalize()
		ctx.Status(http.StatusOK)
	}
}

func init() {
	server.AddHandler("/webui/user", "POST", "/add_subscription_coin", addSubscriptionCoin)
}
