package response

type WebUISuitEntry struct {
	Id         int32  `json:"suit_id"`
	Name       string `json:"suit_name"`
	MemberId   int32  `json:"member_id"`
	MemberName string `json:"member_name"`
	GroupId    int32  `json:"group_id"`
	GroupName  string `json:"group_name"`
	Owned      bool   `json:"owned"`
}

type WebUISuitListResponse []WebUISuitEntry
