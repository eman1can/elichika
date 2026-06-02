package event

import (
	"elichika/internal/client/response"
	"elichika/internal/enum"
	"elichika/internal/subsystem/user_event/mining"
	"elichika/internal/userdata"
)

func FetchEventMining(session *userdata.Session, eventId int32) (*response.FetchEventMiningResponse, *response.RecoverableExceptionResponse) {
	active := session.Gamedata.EventActive
	if (active == nil) || (active.EventId != eventId) {
		return nil, &response.RecoverableExceptionResponse{
			RecoverableExceptionType: enum.RecoverableExceptionTypeEventMiningOutOfDate,
		}
	}

	event := session.Gamedata.EventMining[active.EventId]
	resp := &response.FetchEventMiningResponse{
		EventMiningTopStatus: event.TopStatus,
		UserModelDiff:        &session.UserModel,
	}
	resp.EventMiningTopStatus.StartAt = active.StartAt
	resp.EventMiningTopStatus.EndAt = active.EndAt
	resp.EventMiningTopStatus.ResultAt = active.ResultAt
	resp.EventMiningTopStatus.ExpiredAt = active.ExpiredAt
	userEventStatus := GetUserEventStatus(session, eventId)
	if userEventStatus.IsNew {
		resp.EventMiningTopStatus.IsAddNewPanel = true
		userEventStatus.IsNew = false
		UpdateUserEventStatus(session, userEventStatus)
	}
	userEventMining := mining.GetUserEventMining(session)
	// the engine can only display at most 2 animations, even if it looks like it should display all 3
	// the way it work is something like this:
	// - 1 animation is always unlocked for the gacha characters
	//   - it's pretty much random which gacha characters is shown
	//   - we just choose random here
	// - the animation is unlocked after all story are cleared and is always reserved for the event character
	// - one animation has highest priority, another animation has the lowest priority (after the end)
	// - we have to dynamically set these values pretty much
	if event.HasAnimation() {
		gachaCharacterAnimation := event.RandomGachaCharacterAnimation()
		if userEventMining.ReadStoryNumber < 7 {
			gachaCharacterAnimation.Priority = 160
			resp.EventMiningTopStatus.EventMiningTopAnimationCellMasterRows.Append(gachaCharacterAnimation)
		} else {
			eventCharacterAnimation := event.EventCharacterAnimation()
			gachaCharacterAnimation.Priority = 260
			eventCharacterAnimation.Priority = 210
			resp.EventMiningTopStatus.EventMiningTopAnimationCellMasterRows.Append(eventCharacterAnimation)
			resp.EventMiningTopStatus.EventMiningTopAnimationCellMasterRows.Append(gachaCharacterAnimation)
		}
	}
	n := 0
	for ; n < len(resp.EventMiningTopStatus.EventMiningTopStillCellMasterRows.Slice); n++ {
		if resp.EventMiningTopStatus.EventMiningTopStillCellMasterRows.Slice[n].AddStoryNumber > userEventMining.ReadStoryNumber {
			break
		}
	}
	resp.EventMiningTopStatus.EventMiningTopStillCellMasterRows.Slice = resp.EventMiningTopStatus.EventMiningTopStillCellMasterRows.Slice[:n]
	resp.EventMiningTopStatus.StoryStatus.ReadStoryNumber = userEventMining.ReadStoryNumber
	resp.EventMiningTopStatus.UserRankingStatus = mining.GetEventMiningUserRanking(session)
	resp.EventMiningTopStatus.EventMiningTopCellStateList = mining.GetEventMiningTopCellStateList(session, eventId)
	resp.EventMiningTopStatus.EventMiningMusicBestScoreList = mining.GetUserVoltageRankingSummary(session, eventId).GetMusicBestScoreList()
	return resp, nil
}
