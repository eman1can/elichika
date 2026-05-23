package user_new_badge

import (
	"elichika/internal/client"
	"elichika/internal/generic"
	"elichika/internal/userdata/database"
)

func init() {
	database.AddTable("u_new_badge", generic.UserIdWrapper[client.BootstrapNewBadge]{})
}
