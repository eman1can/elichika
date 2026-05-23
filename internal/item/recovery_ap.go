package item

import (
	"elichika/internal/client"
	"elichika/internal/enum"
)

var (
	TrainingTicket = client.Content{
		ContentType:   enum.ContentTypeRecoveryAp,
		ContentId:     2200,
		ContentAmount: 1,
	}
)
