package event

import (
	"elichika/internal/client/response"
	"elichika/internal/subsystem/user_event/mining"
	"elichika/internal/userdata"
)

// note that thumbnail cell id is unique, event id is provided for easy look up (only send the like data for each event)
// note that there's no unlike, once you like a cell it stay liked, so we don't have to handle the unlike case
func LikeEventMiningPanel(session *userdata.Session, eventId, thumbnailCellId int32) *response.LikeEventMiningPanelResponse {
	mining.UpdateUserEventMiningTopCellState(session, eventId, thumbnailCellId)
	// we need to return the fill list here, otherwise the client will not set like on other cells
	return &response.LikeEventMiningPanelResponse{
		EventMiningTopCellStateList: mining.GetEventMiningTopCellStateList(session, eventId),
	}
}
