package response

type WebUIItem struct {
	Name        string `json:"name"`
	ContentType int32  `json:"content_type"`
	ContentId   int32  `json:"content_id"`
}

type WebUIItemListResponse struct {
	Items []WebUIItem `json:"items"`
}
