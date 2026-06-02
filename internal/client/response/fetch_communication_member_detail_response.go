package response

import (
	"elichika/internal/client"
	"elichika/internal/generic"
)

type FetchCommunicationMemberDetailResponse struct {
	MemberLovePanels generic.Array[client.MemberLovePanelList] `json:"member_love_panels"`
	WeekdayState     client.WeekdayState                       `json:"weekday_state"`
}
