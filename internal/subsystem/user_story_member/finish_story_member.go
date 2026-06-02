package user_story_member

import (
	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/generic"
	"elichika/internal/subsystem/user_info_trigger"
	"elichika/internal/subsystem/user_live_difficulty"
	"elichika/internal/subsystem/user_present"
	"elichika/internal/userdata"
)

func FinishStoryMember(session *userdata.Session, storyMemberMasterId int32) {
	storyMemberMaster := session.Gamedata.StoryMember[storyMemberMasterId]

	userStoryMember := GetStoryMember(session, storyMemberMasterId)

	if userStoryMember.IsNew {
		userStoryMember.IsNew = false
		if storyMemberMaster.Reward != nil {
			user_present.AddPresent(session, client.PresentItem{
				Content:          *storyMemberMaster.Reward,
				PresentRouteType: enum.PresentRouteTypeStoryMember,
				PresentRouteId:   generic.NewNullable(storyMemberMasterId),
			})
			user_info_trigger.AddTriggerBasic(session, client.UserInfoTriggerBasic{
				InfoTriggerType: enum.InfoTriggerTypeStoryMemberReward,
				ParamInt:        generic.NewNullable(storyMemberMasterId),
			})
		}

		UpdateStoryMember(session, userStoryMember)
	}
	// always try to unlock the live
	if storyMemberMaster.UnlockLiveId != nil {
		user_live_difficulty.UnlockLive(session, *storyMemberMaster.UnlockLiveId)
	}

}
