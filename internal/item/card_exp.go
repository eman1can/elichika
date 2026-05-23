package item

import (
	"elichika/internal/client"
	"elichika/internal/enum"
)

var (
	EXP = client.Content{
		ContentType:   enum.ContentTypeCardExp,
		ContentId:     1100,
		ContentAmount: 1,
	}
)
