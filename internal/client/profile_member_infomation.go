package client

import "elichika/internal/generic"

type ProfileMemberInfomation struct {
	UserMembers generic.Array[ProfileUserMember] `json:"user_members"`
}
