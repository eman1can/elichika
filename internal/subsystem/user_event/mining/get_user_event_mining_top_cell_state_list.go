package mining

import (
	"elichika/internal/client"
	"elichika/internal/generic"
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

func GetEventMiningTopCellStateList(session *userdata.Session, eventId int32) generic.List[client.EventMiningTopCellState] {
	eventMiningTopCellStateList := generic.List[client.EventMiningTopCellState]{}
	err := session.Db.Table("u_event_mining_top_cell_state").
		Where("user_id = ? AND event_id = ?", session.UserId, eventId).OrderBy("thumbnail_cell_id").Find(&eventMiningTopCellStateList.Slice)
	utils.CheckErr(err)
	return eventMiningTopCellStateList
}
