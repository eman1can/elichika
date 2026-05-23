package user_live

import (
	"elichika/internal/generic"
	"elichika/internal/subsystem/user_status"
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

func SurrenderLive(session *userdata.Session) generic.Nullable[int32] {
	exist, _, startReq := LoadUserLive(session)
	utils.MustExist(exist)
	ClearUserLive(session)
	// remove only half the LP
	lpCost := session.Gamedata.LiveDifficulty[startReq.LiveDifficultyId].ConsumedLP / 2
	user_status.AddUserLp(session, -lpCost)
	return generic.NewNullable(lpCost)
}
