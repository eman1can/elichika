package user_status

import (
	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/subsystem/user_content"
	"elichika/internal/userdata"
)

// TODO(sns_coin): Handle paid(?)
func addSnsCoin(session *userdata.Session, content *client.Content) any {
	user_content.OverflowCheckedAdd(&session.UserStatus.FreeSnsCoin, &content.ContentAmount)
	return nil
}

func init() {
	user_content.AddContentHandler(enum.ContentTypeSnsCoin, addSnsCoin)
}
