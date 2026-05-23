package user_bootstrap

import (
	"log"

	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/enum"
	"elichika/internal/generic"
	"elichika/internal/subsystem/banner"
	"elichika/internal/subsystem/pickup_info"
	"elichika/internal/subsystem/user_beginner_challenge"
	"elichika/internal/subsystem/user_event/marathon"
	"elichika/internal/subsystem/user_event/mining"
	"elichika/internal/subsystem/user_expired_item"
	"elichika/internal/subsystem/user_login_bonus"
	"elichika/internal/subsystem/user_member_guild"
	"elichika/internal/subsystem/user_new_badge"
	"elichika/internal/subsystem/user_subscription_status"
	"elichika/internal/userdata"
)

func FetchBootstrap(session *userdata.Session, req request.FetchBootstrapRequest) response.FetchBootstrapResponse {
	session.UserStatus.BootstrapSifidCheckAt = session.Time.UnixMilli()
	session.UserStatus.DeviceToken = req.DeviceToken
	// session key will take care of different devices

	resp := response.FetchBootstrapResponse{
		UserModelDiff: &session.UserModel,
	}

	for _, fetchType := range req.BootstrapFetchTypes.Slice {
		switch fetchType {
		case enum.BootstrapFetchTypeBanner:
			resp.FetchBootstrapBannerResponse = banner.GetBootstrapBannerResponse(session)
		case enum.BootstrapFetchTypeNewBadge:
			resp.FetchBootstrapNewBadgeResponse = generic.NewNullable(user_new_badge.GetBootstrapNewBadgeResponse(session))
		case enum.BootstrapFetchTypePickupInfo:
			resp.FetchBootstrapPickupInfoResponse = pickup_info.GetBootstrapPickupInfo(session)
		case enum.BootstrapFetchTypeExpiredItem:
			resp.FetchBootstrapExpiredItemResponse = user_expired_item.GetBootstrapExpiredItem(session)
		case enum.BootstrapFetchTypeGachaPointExchange:
			// TODO(gacha): Implement gacha point
			continue
		case enum.BootstrapFetchTypeLoginBonus:
			resp.FetchBootstrapLoginBonusResponse = generic.NewNullable(user_login_bonus.GetBootstrapLoginBonus(session))
		case enum.BootstrapFetchTypeNotice:
			// TODO(notice): Implement notice
			continue
		case enum.BootstrapFetchTypeSchoolIdolFestivalIdReward:
			// this is no longer used, the client can't comprehend it unless we change stuff
			continue
		default:
			log.Panic("unexpected type")
		}
	}

	// MissionBeginnerMasterId seems to no longer be used, and it's not necessary for the newest set
	resp.ShowChallengeBeginnerButton, resp.ChallengeBeginnerCompletedIds =
		user_beginner_challenge.GetUserBeginnerChallengeBootstrap(session)

	status := user_subscription_status.GetUserSubsriptionStatus(session, 13001)
	session.UserModel.UserSubscriptionStatusById.Set(status.SubscriptionMasterId, status)

	user_member_guild.FetchUserInfoTriggerMemberGuildRankingShowResultRows(session,
		&resp.UserInfoTrigger.UserInfoTriggerMemberGuildRankingShowResultRows)

	marathon.FetchUserInfoTriggerEventMarathonShowResultRows(session,
		&resp.UserInfoTrigger.UserInfoTriggerEventMarathonShowResultRows)

	mining.FetchUserInfoTriggerEventMiningShowResultRows(session,
		&resp.UserInfoTrigger.UserInfoTriggerEventMiningShowResultRows)
	return resp
}
