package response

import "elichika/internal/client"

type FetchNoticeDetailResponse struct {
	Notice client.NoticeDetail `json:"notice"`
}
