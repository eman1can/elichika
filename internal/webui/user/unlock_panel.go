package user

import (
	"net/http"
	"slices"

	"elichika/internal/server"
	"elichika/internal/subsystem/user_member"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

type WebUIUnlockLovePanelRequest struct {
	MemberMasterIds []int32 `form:"member_master_ids" json:"member_master_ids"`
}

func unlockLovePanel(ctx *gin.Context) {
	req := WebUIUnlockLovePanelRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := ctx.MustGet("session").(*userdata.Session)

	for panelId, masterPanel := range session.Gamedata.MemberLovePanel {
		if !slices.Contains(req.MemberMasterIds, *masterPanel.MemberMasterId) {
			continue
		}

		lovePanel := user_member.GetMemberLovePanel(session, *masterPanel.MemberMasterId, panelId)
		lovePanel.Status = 0x1F // b00011111
		user_member.UpdateMemberLovePanel(session, lovePanel)
	}

	session.Finalize()
	ctx.JSON(http.StatusOK, gin.H{})
}

func init() {
	server.AddHandler("/webui/user", "POST", "/unlock_love_panel", unlockLovePanel)
}
