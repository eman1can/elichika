package response

import (
	"elichika/internal/client"
	"elichika/internal/generic"
)

type FetchShopPackResponse struct {
	ProductList      generic.Array[client.ShopBillingProduct] `json:"product_list"`
	BillingStateInfo client.BillingStateInfo                  `json:"billing_state_info"`
	InPeriodGiftBox  generic.Array[client.InPeriodGiftBox]    `json:"in_period_gift_box"`
}
