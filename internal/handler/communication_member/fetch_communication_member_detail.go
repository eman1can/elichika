package communication_member

import (
	"encoding/json"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/time"
	"elichika/internal/subsystem/user_member"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func fetchCommunicationMemberDetail(ctx *gin.Context) {
	req := request.FetchCommunicationMemberDetailRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	resp := response.FetchCommunicationMemberDetailResponse{}
	resp.MemberLovePanels.Append(user_member.GetMemberLovePanel(session, req.MemberId))

	resp.WeekdayState = time.GetWeekdayState(session)
	common.JsonResponse(ctx, resp)
}

func init() {
	server.AddHandler("/", "POST", "/communicationMember/fetchCommunicationMemberDetail", fetchCommunicationMemberDetail)
}
