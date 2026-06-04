package user_status

import (
	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/subsystem/user_content"
	"elichika/internal/userdata"
)

// TODO(sns_coin): Handle paid(?)
func addSnsCoin(session *userdata.Session, content *client.Content) any {
	if content.ContentType == enum.ContentTypeSnsCoin {
		if content.ContentId == enum.SnsCoinFree {
			user_content.OverflowCheckedAdd(&session.UserStatus.FreeSnsCoin, &content.ContentAmount)
		} else if content.ContentId == enum.SnsCoinGoogle {
			user_content.OverflowCheckedAdd(&session.UserStatus.GoogleSnsCoin, &content.ContentAmount)
		} else {
			user_content.OverflowCheckedAdd(&session.UserStatus.AppleSnsCoin, &content.ContentAmount)
		}
	}
	return nil
}

func init() {
	user_content.AddContentHandler(enum.ContentTypeSnsCoin, addSnsCoin)
}
