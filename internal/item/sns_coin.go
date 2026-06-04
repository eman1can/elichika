package item

import (
	"elichika/internal/client"
	"elichika/internal/enum"
)

var (
	StarGem = client.Content{
		ContentType:   enum.ContentTypeSnsCoin,
		ContentId:     enum.SnsCoinFree,
		ContentAmount: 1,
	}
	StarGemApple = client.Content{
		ContentType:   enum.ContentTypeSnsCoin,
		ContentId:     enum.SnsCoinApple,
		ContentAmount: 1,
	}
	StarGemGoogle = client.Content{
		ContentType:   enum.ContentTypeSnsCoin,
		ContentId:     enum.SnsCoinGoogle,
		ContentAmount: 1,
	}
)
