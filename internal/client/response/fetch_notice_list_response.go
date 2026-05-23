package response

import "elichika/internal/client"

type FetchNoticeListResponse struct {
	NoticeList client.NoticeList `json:"notice_list"`
}
