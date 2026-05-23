package item

import (
	"elichika/internal/client"
	"elichika/internal/enum"
)

var (
	ShowCandy50 = client.Content{
		ContentType:   enum.ContentTypeRecoveryLp,
		ContentId:     1300,
		ContentAmount: 1,
	}
	ShowCandy100 = client.Content{
		ContentType:   enum.ContentTypeRecoveryLp,
		ContentId:     1301,
		ContentAmount: 1,
	}
)
