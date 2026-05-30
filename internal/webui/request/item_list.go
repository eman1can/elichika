package request

type WebUIItemListRequest struct {
	ContentType int32  `form:"type"`
	Language    string `form:"l"`
}
