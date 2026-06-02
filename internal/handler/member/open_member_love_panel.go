package member

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_member"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func openMemberLovePanel(ctx *gin.Context) {
	req := request.OpenMemberLovePanelRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	resp := response.OpenMemberLovePanelResponse{
		UserModel: &session.UserModel,
	}

	resp.MemberLovePanels.Append(user_member.UnlockMemberLovePanelCells(session, req.MemberId, req.MemberLovePanelId, req.MemberLovePanelCellIds.Slice))

	common.JsonResponse(ctx, &resp)
}

func init() {
	server.AddHandler("/", "POST", "/member/openMemberLovePanel", openMemberLovePanel)
}
