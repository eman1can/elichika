package response

type WebUIBackgroundEntry struct {
	Id           int32  `json:"id"`
	Name         string `json:"name"`
	DisplayOrder int32  `json:"display_order"`
	Owned        bool   `json:"owned"`
}
type WebUIBackgroundListResponse []WebUIBackgroundEntry
