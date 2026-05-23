package client

import "elichika/internal/generic"

type PartnerSelectState struct {
	LivePartners generic.Array[LivePartner] `json:"live_partners"`
	FriendCount  int32                      `json:"friend_count"`
}
