package request

import (
	"elichika/internal/client"
	"elichika/internal/generic"
)

type UpdateGdprConsentStateRequest struct {
	Version     int32                                `json:"version"`
	ConsentList generic.List[client.GdprConsentInfo] `json:"consent_list"`
}
