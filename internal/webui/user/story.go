package user

type WebUIStoryCellEntry struct {
	Id             int32  `json:"id"`
	Chapter        int32  `json:"chapter"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	ImageAssetPath string `json:"image_asset_path"`
	IsNew          bool   `json:"is_new"`
	Unlocked       bool   `json:"unlocked"`
}

type WebUIStoryChapterEntry struct {
	Id             int32                 `json:"id"`
	Title          string                `json:"title"`
	DisplayOrder   int32                 `json:"display_order"`
	Description    string                `json:"description"`
	ImageAssetPath string                `json:"image_asset_path"`
	Chapters       []WebUIStoryCellEntry `json:"chapters"`
}
