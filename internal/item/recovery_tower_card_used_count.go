package item

import (
	"elichika/internal/client"
	"elichika/internal/enum"
)

var (
	PerformanceDrink = client.Content{
		ContentType:   enum.ContentTypeRecoveryTowerCardUsedCount,
		ContentId:     24001,
		ContentAmount: 1,
	}
)
