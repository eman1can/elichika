package voltage_ranking

import (
	"elichika/internal/client"
	"elichika/internal/client/response"
	"elichika/internal/subsystem/cache"
	"elichika/internal/subsystem/user_social"
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

var (
	getVoltageRankingDeckResponseCache = cache.UniquePointerMap[int64, cache.CachedObject[client.OtherUserDeckDetail]]{}
)

func GetVoltageRankingDeckResponse(session *userdata.Session, liveDifficultyId int32, userId int32) response.GetVoltageRankingDeckResponse {
	key := (int64(liveDifficultyId) << 32) ^ (int64(userId))
	cacher := getVoltageRankingDeckResponseCache.Get(key)
	cacher.Acquire()
	defer cacher.Release()
	if cacher.ExpireAt <= session.Time.Unix() {
		cacher.ExpireAt = session.Time.Unix() + VoltageRankingDeckCache
		cacher.Value = getVoltageRankingDeckNoCache(session, liveDifficultyId, userId)
	}
	return response.GetVoltageRankingDeckResponse{
		User:       user_social.GetOtherUser(session, userId),
		DeckDetail: *cacher.Value,
	}
}

func getVoltageRankingDeckNoCache(session *userdata.Session, liveDifficultyId int32, userId int32) *client.OtherUserDeckDetail {
	resp := response.GetVoltageRankingDeckResponse{}
	exist, err := session.Db.Table("u_voltage_ranking").
		Where("live_difficulty_id = ? AND user_id = ?", liveDifficultyId, userId).
		Cols("deck_detail").Get(&resp) // can't get directly into an object because weird xorm quirk
	utils.CheckErrMustExist(err, exist)
	return &resp.DeckDetail
}
