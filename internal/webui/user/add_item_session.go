package user

import (
	"net/http"

	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

type WebUIAddItemRequest struct {
	Amount int32 `form:"amount" json:"amount"`
}

func getAddItemSession(ctx *gin.Context) (*userdata.Session, int32) {
	req := WebUIAddItemRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, 0
	}

	session := ctx.MustGet("session").(*userdata.Session)

	return session, req.Amount
}
