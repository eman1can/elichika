package response

type WebUIMemberEntry struct {
	Id                   int32  `json:"member_id"`
	Name                 string `json:"member_name"`
	GroupId              int32  `json:"group_id"`
	GroupName            string `json:"group_name"`
	RepresentativeCardId int32  `json:"card_id"`
	LoveLevel            int32  `json:"love_level"`
}

type WebUIMemberListResponse []WebUIMemberEntry
