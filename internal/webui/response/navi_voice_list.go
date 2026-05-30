package response

type WebUINaviVoiceEntry struct {
	Id           int32  `json:"id"`
	Name         string `json:"name"`
	MemberId     int32  `json:"member_id"`
	MemberName   string `json:"member_name"`
	GroupId      int32  `json:"group_id"`
	GroupName    string `json:"group_name"`
	DisplayOrder int32  `json:"display_order"`
	ReleaseRoute int32  `json:"release_route"`
	ReleaseValue int32  `json:"release_value"`
	Owned        bool   `json:"owned"`
}

type WebUINaviVoiceListResponse []WebUINaviVoiceEntry
