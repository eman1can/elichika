package user

import (
	"net/http"

	"elichika/internal/server"
	"elichika/internal/subsystem/user_suit"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

type WebUIAddSuitRequest struct {
	SuitMasterIds []int32 `form:"suit_master_ids" json:"suit_master_ids"`
}

func addSuit(ctx *gin.Context) {
	var req WebUIAddSuitRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := ctx.MustGet("session").(*userdata.Session)

	for _, suitMasterId := range req.SuitMasterIds {
		if !user_suit.HasSuit(session, suitMasterId) {
			user_suit.InsertUserSuit(session, suitMasterId)
		}
	}

	session.Finalize()
	ctx.JSON(http.StatusOK, gin.H{})
}

func init() {
	server.AddHandler("/webui/user", "POST", "/add_suit", addSuit)
}
