package response

import (
	"elichika/internal/client"
	"elichika/internal/generic"
)

type UserSearchResponse struct {
	UserSearchList generic.Array[client.OtherUser] `json:"user_search_list"`
}
