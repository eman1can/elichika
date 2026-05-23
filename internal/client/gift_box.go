package client

import "elichika/internal/generic"

type GiftBox struct {
	IsInPeriodGiftBox bool                          `json:"is_in_period_gift_box"`
	GiftBoxContent    generic.Array[GiftBoxContent] `json:"gift_box_content"`
}
