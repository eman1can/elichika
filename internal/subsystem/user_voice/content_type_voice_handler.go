package user_voice

import (
	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/subsystem/user_content"
	"elichika/internal/userdata"
)

func contentTypeVoiceHandler(session *userdata.Session, content *client.Content) any {
	UpdateUserVoice(session, content.ContentId, false)
	content.ContentAmount = 0
	return nil
}

func init() {
	user_content.AddContentHandler(enum.ContentTypeVoice, contentTypeVoiceHandler)
}
