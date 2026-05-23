package response

import "elichika/internal/client"

type SetTakeOverResponse struct {
	Data client.UserLinkData `json:"data"`
}
