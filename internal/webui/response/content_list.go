package response

type WebUIContent struct {
	Name        string `json:"name"`
	ContentType int32  `json:"content_type"`
}

type WebUIContentListResponse struct {
	Items []WebUIContent `json:"items"`
}
