package user_social

import (
	"elichika/internal/client/response"
	"elichika/internal/enum"
	"elichika/internal/userdata"
)

func RemoveFriendOtherScene(session *userdata.Session, otherUserId int32) (*response.FriendActionResponse, *response.FriendActionRecoverableExceptionResponse) {
	RemoveConnection(session, otherUserId)
	return &response.FriendActionResponse{
		SuccessType:  enum.FriendSuccessTypeNoProblem,
		TargetPlayer: GetNullableOtherUser(session, otherUserId),
	}, nil
}
