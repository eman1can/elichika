package request

type WebUIItemImageRequest struct {
	ContentType int32 `form:"type"`
	ContentId   int32 `form:"id"`
}
