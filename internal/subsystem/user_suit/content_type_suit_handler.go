package user_suit

import (
	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/subsystem/user_content"
	"elichika/internal/userdata"
)

func contentTypeSuitHandler(session *userdata.Session, content *client.Content) any {
	InsertUserSuit(session, content.ContentId)
	content.ContentAmount = 0
	return nil
}

func init() {
	user_content.AddContentHandler(enum.ContentTypeSuit, contentTypeSuitHandler)
}
