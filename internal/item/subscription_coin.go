package item

import (
	"elichika/internal/client"
	"elichika/internal/enum"
)

var (
	MemberCoin = client.Content{
		ContentType:   enum.ContentTypeSubscriptionCoin,
		ContentId:     0,
		ContentAmount: 1,
	}
)
