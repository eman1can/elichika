package response

type WebUICardEntry struct {
	Id                     int32  `json:"card_id"`
	Name                   string `json:"card_name"`
	MemberId               int32  `json:"member_id"`
	MemberName             string `json:"member_name"`
	GroupId                int32  `json:"group_id"`
	GroupName              string `json:"group_name"`
	Grade                  int32  `json:"grade"`
	Rarity                 int32  `json:"rarity"`
	IsAllTrainingActivated bool   `json:"is_all_training_activated"`
}

type WebUICardListResponse []WebUICardEntry
