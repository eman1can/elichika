package client

import "elichika/internal/generic"

type BootstrapExpiredItem struct {
	ExpiredItems generic.Array[Content] `json:"expired_items"`
}
