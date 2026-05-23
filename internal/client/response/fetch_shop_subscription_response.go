package response

import (
	"elichika/internal/client"
	"elichika/internal/generic"
)

type FetchShopSubscriptionResponse struct {
	ProductList      generic.Array[client.ShopBillingProduct] `json:"product_list"`
	BillingStateInfo client.BillingStateInfo                  `json:"billing_state_info"`
}
