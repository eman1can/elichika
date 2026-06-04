package user_status

import (
	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/subsystem/user_content"
	"elichika/internal/userdata"
)

func addGameMoney(session *userdata.Session, content *client.Content) any {
	if content.ContentType == enum.ContentTypeGameMoney {
		user_content.OverflowCheckedAdd(&session.UserStatus.GameMoney, &content.ContentAmount)
	}
	return nil
}

func init() {
	user_content.AddContentHandler(enum.ContentTypeGameMoney, addGameMoney)
}
