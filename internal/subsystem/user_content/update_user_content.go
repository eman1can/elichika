package user_content

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
)

func UpdateUserContent(session *userdata.Session, content client.Content) {
	_, exist := session.UserContentDiffs[content.ContentType]
	if !exist {
		session.UserContentDiffs[content.ContentType] = make(map[int32]client.Content)
	}
	session.UserContentDiffs[content.ContentType][content.ContentId] = content
}
