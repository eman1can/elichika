package item

import (
	"elichika/internal/client"
	"elichika/internal/enum"
)

func MementoPiece(memberId int32) *client.Content {
	return &client.Content{
		ContentType:   enum.ContentTypeTrainingMaterial,
		ContentId:     int32(18000 + memberId),
		ContentAmount: 1,
	}
}

func MemorialPiece(memberId int32) *client.Content {
	return &client.Content{
		ContentType:   enum.ContentTypeTrainingMaterial,
		ContentId:     int32(8000 + memberId),
		ContentAmount: 1,
	}
}

var (
	IdolBadge = client.Content{
		ContentType:   enum.ContentTypeTrainingMaterial,
		ContentId:     1200,
		ContentAmount: 1,
	}
)
