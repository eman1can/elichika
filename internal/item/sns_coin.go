package item

import (
	"elichika/internal/client"
	"elichika/internal/enum"
)

var (
	StarGem = client.Content{
		ContentType:   enum.ContentTypeSnsCoin,
		ContentId:     0,
		ContentAmount: 1,
	}
)
