package user_content

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
)

func RemoveContent(session *userdata.Session, content client.Content) any {
	content.ContentAmount = -content.ContentAmount
	return AddContent(session, content)
}
