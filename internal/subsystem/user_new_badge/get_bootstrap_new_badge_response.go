package user_new_badge

import (
	"elichika/internal/client"
	"elichika/internal/subsystem/user_mission"
	"elichika/internal/subsystem/user_present"
	"elichika/internal/subsystem/user_social"
	"elichika/internal/userdata"
)

func GetBootstrapNewBadgeResponse(session *userdata.Session) client.BootstrapNewBadge {
	return client.BootstrapNewBadge{
		IsNewMainStory:                     false,
		UnreceivedPresentBox:               user_present.FetchPresentCount(session),
		IsUnreceivedPresentBoxSubscription: false, // TODO(present box, subscription)
		IsUpdateFriend:                     user_social.IsUpdateFriend(session),
		UnreceivedMission:                  user_mission.CountUnreceivedMission(session),
		UnreceivedChallengeBeginner:        0, // TODO(beginner guide)
	}
}
