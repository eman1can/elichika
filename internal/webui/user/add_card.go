package user

import (
	"net/http"

	"elichika/internal/server"
	"elichika/internal/subsystem/user_card"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

type WebUIAddCardRequest struct {
	CardMasterIds []int32 `form:"card_master_ids" json:"card_master_ids"`
}

func addCard(ctx *gin.Context) {
	var req WebUIAddCardRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := ctx.MustGet("session").(*userdata.Session)

	for _, cardMasterId := range req.CardMasterIds {
		user_card.AddUserCardByCardMasterId(session, cardMasterId)
	}

	session.Finalize()
	ctx.JSON(http.StatusOK, gin.H{})
}

func init() {
	server.AddHandler("/webui/user", "POST", "/add_card", addCard)
}
