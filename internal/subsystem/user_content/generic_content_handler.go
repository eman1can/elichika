package user_content

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
)

func genericContentHandler(session *userdata.Session, addedContent *client.Content) any {
	currentContent := GetUserContent(session, addedContent.ContentType, addedContent.ContentId)
	OverflowCheckedAdd(&currentContent.ContentAmount, &addedContent.ContentAmount)
	UpdateUserContent(session, currentContent)
	return nil
}
