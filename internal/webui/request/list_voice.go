package request

type WebUIListVoiceRequest struct {
	Language     string `form:"l"`
	ReleaseRoute *int32 `form:"route"`
	ListType     *int32 `form:"list_type"`
}
