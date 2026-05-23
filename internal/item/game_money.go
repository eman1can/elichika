package item

import (
	"elichika/internal/client"
	"elichika/internal/enum"
)

var (
	Gold = client.Content{
		ContentType:   enum.ContentTypeGameMoney,
		ContentId:     1200,
		ContentAmount: 1,
	}
)
