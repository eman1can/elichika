package shop

import (
	"elichika/internal/client"
	"elichika/internal/client/response"
	"elichika/internal/enum"
	"elichika/internal/generic"
	"elichika/internal/handler/common"
	"elichika/internal/server"

	"github.com/gin-gonic/gin"
)

// Not implementing this system seriously

func fetchShopTop(ctx *gin.Context) {
	// There is no request body
	resp := response.FetchShopTopResponse{}
	{
		arr := generic.Array[client.ShopTopIsOpen]{}
		arr.Append(client.ShopTopIsOpen{
			IsOpen: true,
		})
		resp.IsOpenByShopType.Set(enum.ShopTypeBillingPack, arr)
	}
	{
		arr := generic.Array[client.ShopTopIsOpen]{}
		arr.Append(client.ShopTopIsOpen{
			IsOpen: true,
		})
		resp.IsOpenByShopType.Set(enum.ShopTypeBillingNormal, arr)
	}
	{
		arr := generic.Array[client.ShopTopIsOpen]{}
		arr.Append(client.ShopTopIsOpen{
			IsOpen: false,
		})
		resp.IsOpenByShopType.Set(enum.ShopTypeEventExchange, arr)
	}
	{
		arr := generic.Array[client.ShopTopIsOpen]{}
		arr.Append(client.ShopTopIsOpen{
			IsOpen: false,
		})
		resp.IsOpenByShopType.Set(enum.ShopTypeItemExchange, arr)
	}
	{
		arr := generic.Array[client.ShopTopIsOpen]{}
		arr.Append(client.ShopTopIsOpen{
			IsOpen: false,
		})
		resp.IsOpenByShopType.Set(enum.ShopTypeBillingSubscription, arr)
	}
	common.JsonResponse(ctx, &resp)
}

func init() {
	server.AddHandler("/", "POST", "/shop/fetchShopTop", fetchShopTop)
}
