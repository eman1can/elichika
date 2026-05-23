package user_social

import (
	"elichika/internal/client/response"
	"elichika/internal/enum"
	"elichika/internal/userdata"
)

// cancel a friend request from other menus (not friend menu)
func CancelFriendRequestOtherScene(session *userdata.Session, otherUserId int32) (*response.FriendActionResponse, *response.FriendActionRecoverableExceptionResponse) {
	good, errorKey := CancelFriendRequestShared(session, otherUserId)
	if good {
		return &response.FriendActionResponse{
			SuccessType:  enum.FriendSuccessTypeNoProblem,
			TargetPlayer: GetNullableOtherUser(session, otherUserId),
		}, nil
	} else {
		return nil, &response.FriendActionRecoverableExceptionResponse{
			ErrorKey:     errorKey,
			TargetPlayer: GetNullableOtherUser(session, otherUserId),
		}
	}
}
