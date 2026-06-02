package event

import (
	"elichika/internal/client"
	"elichika/internal/client/response"
	"elichika/internal/enum"
	"elichika/internal/generic"
	"elichika/internal/subsystem/user_event/marathon"
	"elichika/internal/subsystem/user_info_trigger"
	"elichika/internal/userdata"
)

func FetchEventMarathon(session *userdata.Session, eventId int32) (*response.FetchEventMarathonResponse, *response.RecoverableExceptionResponse) {
	active := session.Gamedata.EventActive
	if (active == nil) || (active.EventId != eventId) {
		return nil, &response.RecoverableExceptionResponse{
			RecoverableExceptionType: enum.RecoverableExceptionTypeEventMarathonOutOfDate,
		}
	}

	event := session.Gamedata.EventMarathon[eventId]
	resp := &response.FetchEventMarathonResponse{
		EventMarathonTopStatus: event.TopStatus,
		UserModelDiff:          &session.UserModel,
	}
	resp.EventMarathonTopStatus.StartAt = active.StartAt
	resp.EventMarathonTopStatus.EndAt = active.EndAt
	resp.EventMarathonTopStatus.ResultAt = active.ResultAt
	resp.EventMarathonTopStatus.ExpiredAt = active.ExpiredAt
	userEventStatus := GetUserEventStatus(session, eventId)

	resp.EventMarathonTopStatus.IsFirstAccess = userEventStatus.IsFirstAccess

	if resp.EventMarathonTopStatus.IsFirstAccess {
		user_info_trigger.AddTriggerBasic(session,
			client.UserInfoTriggerBasic{
				InfoTriggerType: enum.InfoTriggerTypeEventMarathonFirstRuleDescription,
				ParamInt:        generic.NewNullable(eventId),
			})
	}

	userEventMarathon := marathon.GetUserEventMarathon(session)

	switch userEventMarathon.ReadStoryNumber {
	case 0:
		resp.EventMarathonTopStatus.BoardStatus.BoardThingMasterRows.Append(event.BoardMemos[0])
	case 7:
		resp.EventMarathonTopStatus.BoardStatus.BoardThingMasterRows.Append(event.BoardMemos[2])
	default:
		resp.EventMarathonTopStatus.BoardStatus.BoardThingMasterRows.Append(event.BoardMemos[1])
	}

	resp.EventMarathonTopStatus.StoryStatus.ReadStoryNumber = userEventMarathon.ReadStoryNumber
	for i := int32(0); i < userEventMarathon.ReadStoryNumber; i++ {
		resp.EventMarathonTopStatus.BoardStatus.BoardThingMasterRows.Append(event.BoardPictures[i])
	}
	if userEventMarathon.ReadStoryNumber > 0 {
		resp.EventMarathonTopStatus.BoardStatus.BoardThingMasterRows.Slice[userEventMarathon.ReadStoryNumber].IsEffect = userEventStatus.IsNew
		resp.EventMarathonTopStatus.BoardStatus.IsEffect = userEventStatus.IsNew
	}
	if resp.EventMarathonTopStatus.BoardStatus.IsEffect || resp.EventMarathonTopStatus.IsFirstAccess {
		userEventStatus.IsFirstAccess = false
		userEventStatus.IsNew = false
		UpdateUserEventStatus(session, userEventStatus)
	}

	nextRewardPoint, nextRewardContent := event.GetNextReward(session.Gamedata, userEventMarathon.EventPoint)

	resp.EventMarathonTopStatus.UserRankingStatus = client.EventMarathonUserRanking{
		Order:           marathon.GetUserEventMarathonRanking(session, event.EventId),
		TotalPoint:      generic.NewNullable(userEventMarathon.EventPoint),
		NextRewardPoint: nextRewardPoint,
		RewardContent:   nextRewardContent,
	}
	return resp, nil
}
