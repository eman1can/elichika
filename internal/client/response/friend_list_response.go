package response

import "elichika/internal/client"

type FriendListResponse struct {
	SuccessType    int32                 `json:"success_type" enum:"FriendSuccessType"`
	FriendViewList client.FriendViewList `json:"friend_view_list"`
}
