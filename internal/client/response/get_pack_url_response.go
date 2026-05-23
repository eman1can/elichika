package response

import "elichika/internal/generic"

type GetPackUrlResponse struct {
	UrlList generic.List[string] `json:"url_list"`
}
