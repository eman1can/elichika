package user

import (
	"net/http"

	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

type WebUIMasterIdRequest struct {
	MasterIds []int32 `json:"master_ids"`
}

func getMasterIdsSession(ctx *gin.Context) (*userdata.Session, []int32) {
	var req WebUIMasterIdRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, []int32{}
	}

	session := ctx.MustGet("session").(*userdata.Session)

	return session, req.MasterIds
}
