package response

import (
	"elichika/internal/client"
	"elichika/internal/generic"
)

type FetchPresentResponse struct {
	PresentItems        generic.List[client.PresentItem]        `json:"present_items"`
	PresentHistoryItems generic.List[client.PresentHistoryItem] `json:"present_history_items"`
	PresentCount        int32                                   `json:"present_count"`
}
