package response

import (
	"elichika/internal/client"
	"elichika/internal/generic"
)

type FriendRecoverableExceptionResponse struct {
	ErrorKey       int32                                   `json:"error_key" enum:"FriendFailureType"`
	FriendViewList generic.Nullable[client.FriendViewList] `json:"friend_view_list"`
}
