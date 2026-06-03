package user

import (
	"net/http"

	"elichika/internal/server"
	"elichika/internal/subsystem/user_member"
	"elichika/internal/subsystem/user_story_member"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

type WebUIFinishMemberStoryRequest struct {
	MemberMasterIds []int32 `form:"member_master_ids" json:"member_master_ids"`
}

func finishMemberStory(ctx *gin.Context) {
	var req WebUIFinishMemberStoryRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := ctx.MustGet("session").(*userdata.Session)

	for _, memberMasterId := range req.MemberMasterIds {
		for _, storyMember := range session.Gamedata.StoryMember {
			if storyMember.MemberMId != memberMasterId {
				continue
			}

			member := user_member.GetMember(session, memberMasterId)
			if member.LoveLevel < storyMember.LoveLevel {
				continue
			}

			user_story_member.FinishStoryMember(session, storyMember.Id)
		}
	}

	session.UserModel.UserInfoTriggerBasicByTriggerId.Clear()
	session.Finalize()
	ctx.JSON(http.StatusOK, gin.H{})
}

func init() {
	server.AddHandler("/webui/user", "POST", "/finish_member_story", finishMemberStory)
}
