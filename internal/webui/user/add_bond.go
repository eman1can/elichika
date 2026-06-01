package user

import (
	"net/http"

	"elichika/internal/server"
	"elichika/internal/subsystem/user_member"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

type WebUIAddBondRequest struct {
	MemberMasterIds []int32 `form:"member_master_ids" json:"member_master_ids"`
	Amount          int32   `form:"amount" json:"amount"`
}

func addBond(ctx *gin.Context) {
	req := WebUIAddBondRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := ctx.MustGet("session").(*userdata.Session)

	for _, memberMasterId := range req.MemberMasterIds {
		user_member.AddMemberLovePoint(session, memberMasterId, req.Amount)
	}

	session.Finalize()
	ctx.JSON(http.StatusOK, gin.H{})
}

func init() {
	server.AddHandler("/webui/user", "POST", "/add_bond", addBond)
}
