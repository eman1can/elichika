package user_social

import (
	"elichika/internal/client/response"
	"elichika/internal/enum"
	"elichika/internal/generic"
	"elichika/internal/userdata"
)

func SearchUserId(session *userdata.Session, otherUserId int32) (
	*response.UserSearchResponse, *response.FriendRecoverableExceptionResponse) {
	if !UserExist(session, otherUserId) {
		return nil, &response.FriendRecoverableExceptionResponse{
			ErrorKey:       enum.FriendFailureTypeSearchNotExist,
			FriendViewList: generic.NewNullable(GetFriendViewList(session)),
		}
	} else {
		resp := response.UserSearchResponse{}
		resp.UserSearchList.Append(GetOtherUser(session, otherUserId))
		return &resp, nil
	}
}
