package user_member

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
)

func InsertMembers(session *userdata.Session, members []client.UserMember) {
	for _, member := range members {
		UpdateMember(session, member)
	}
}
