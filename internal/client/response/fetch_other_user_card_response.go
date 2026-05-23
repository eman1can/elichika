package response

import "elichika/internal/client"

type FetchOtherUserCardResponse struct {
	OtherUserCard client.OtherUserCard `json:"other_user_card"`
}
