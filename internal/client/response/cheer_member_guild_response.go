package response

import (
	"elichika/internal/client"
	"elichika/internal/generic"
)

type CheerMemberGuildResponse struct {
	Rewards              generic.Array[client.Content] `json:"rewards"`
	MemberGuildTopStatus client.MemberGuildTopStatus   `json:"member_guild_top_status"`
	UserModelDiff        *client.UserModel             `json:"user_model_diff"`
}
