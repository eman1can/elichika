package item

import (
	"elichika/internal/client"
	"elichika/internal/enum"
)

var (
	MemoryKey = client.Content{
		ContentType:   enum.ContentTypeStoryEventUnlock,
		ContentId:     17001,
		ContentAmount: 1,
	}
)
