package communication_member

import (
	"log"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/enum"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_member"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func updateUserCommunicationMemberDetailBadge(ctx *gin.Context) {
	req := request.UpdateMemberDetailBadgeRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	detailBadge := user_member.GetUserCommunicationMemberDetailBadge(session, req.MemberMasterId)
	switch req.CommunicationMemberDetailBadgeType {
	case enum.CommunicationMemberDetailBadgeTypeStoryMember:
		detailBadge.IsStoryMemberBadge = false
	case enum.CommunicationMemberDetailBadgeTypeStorySide:
		detailBadge.IsStorySideBadge = false
	case enum.CommunicationMemberDetailBadgeTypeVoice:
		detailBadge.IsVoiceBadge = false
	case enum.CommunicationMemberDetailBadgeTypeTheme:
		detailBadge.IsThemeBadge = false
	case enum.CommunicationMemberDetailBadgeTypeCard:
		detailBadge.IsCardBadge = false
	case enum.CommunicationMemberDetailBadgeTypeMusic:
		detailBadge.IsMusicBadge = false
	default:
		log.Panic("unknown type")
	}
	user_member.UpdateUserCommunicationMemberDetailBadge(session, detailBadge)

	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	server.AddHandler("/", "POST", "/communicationMember/updateUserCommunicationMemberDetailBadge", updateUserCommunicationMemberDetailBadge)
}
