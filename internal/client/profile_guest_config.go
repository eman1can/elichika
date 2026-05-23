package client

import "elichika/internal/generic"

type ProfileGuestConfig struct {
	LivePartnerCards generic.Array[ProfileLivePartnerCard] `json:"live_partner_cards"`
}
