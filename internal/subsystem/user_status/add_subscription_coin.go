package user_status

import (
	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/subsystem/user_content"
	"elichika/internal/userdata"
)

func addSubscriptionCoin(session *userdata.Session, content *client.Content) any {
	user_content.OverflowCheckedAdd(&session.UserStatus.SubscriptionCoin, &content.ContentAmount)
	return nil
}

func init() {
	user_content.AddContentHandler(enum.ContentTypeSubscriptionCoin, addSubscriptionCoin)
}
