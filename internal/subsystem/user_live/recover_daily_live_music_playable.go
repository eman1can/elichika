package user_live

import (
	"elichika/internal/client/response"
	"elichika/internal/config"
	"elichika/internal/enum"
	"elichika/internal/item"
	"elichika/internal/subsystem/user_content"
	"elichika/internal/userdata"
)

func RecoverDailyLiveMusicPlayable(session *userdata.Session, liveId int32) (*response.RecoverDailyLiveMusicPlayableResponse,
	*response.RecoverableExceptionResponse) {

	liveDailyMasterId := GetLiveDailyMasterId(session, liveId)
	if liveDailyMasterId == nil {
		return nil, &response.RecoverableExceptionResponse{
			RecoverableExceptionType: enum.RecoverableExceptionTypeDailyLiveOutOfTerm,
		}
	}
	userLiveDaily := GetUserLiveDaily(session, *liveDailyMasterId)
	userLiveDaily.RemainingPlayCount = session.Gamedata.ConstantInt[enum.ConstantIntDailyLivePlayLimitRecoverOnce]
	if config.Conf.ResourceConfig().ConsumeMiscItems {
		userLiveDaily.RemainingRecoveryCount.Value--
		user_content.RemoveContent(session, item.StarGem.Amount(session.Gamedata.ConstantInt[enum.ConstantIntDailyLivePlayLimitRecoverCost]))
	}
	UpdateUserLiveDaily(session, userLiveDaily)
	return &response.RecoverDailyLiveMusicPlayableResponse{
		UserModelDiff: &session.UserModel,
		LiveDaily:     userLiveDaily,
	}, nil
}
