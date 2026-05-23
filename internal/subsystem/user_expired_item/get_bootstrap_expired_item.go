package user_expired_item

import (
	"elichika/internal/client"
	"elichika/internal/generic"
	"elichika/internal/userdata"
)

// TODO(present_box): Handle expired items
func GetBootstrapExpiredItem(session *userdata.Session) generic.Nullable[client.BootstrapExpiredItem] {
	return generic.Nullable[client.BootstrapExpiredItem]{}
}
