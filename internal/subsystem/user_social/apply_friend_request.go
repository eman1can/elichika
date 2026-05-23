package user_social

import (
	"elichika/internal/client/response"
	"elichika/internal/enum"
	"elichika/internal/generic"
	"elichika/internal/userdata"
)

// add a friend request from the friend menu
func ApplyFriendRequest(session *userdata.Session, otherUserId int32) (*response.FriendListResponse, *response.FriendRecoverableExceptionResponse) {
	good, errorKey := ApplyFriendRequestShared(session, otherUserId)
	if good {
		return &response.FriendListResponse{
			SuccessType:    enum.FriendSuccessTypeNoProblem,
			FriendViewList: GetFriendViewList(session),
		}, nil
	} else {
		return nil, &response.FriendRecoverableExceptionResponse{
			ErrorKey:       errorKey,
			FriendViewList: generic.NewNullable(GetFriendViewList(session)),
		}
	}
}
