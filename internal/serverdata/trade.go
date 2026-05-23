package serverdata

import (
	"encoding/json"

	"elichika/internal/client"
	"elichika/internal/config"
	utils2 "elichika/internal/utils"

	"xorm.io/xorm"
)

func InsertTrade(session *xorm.Session) {
	// insert trades from json

	file := config.ServerInitJsons + "trade.json"

	trades := []client.Trade{}
	err := json.Unmarshal([]byte(utils2.ReadAllText(file)), &trades)
	utils2.CheckErr(err)

	for i, trade := range trades {
		trades[i].EndAt.HasValue = false
		trades[i].ResetAt.HasValue = false
		for j, product := range trade.Products.Slice {
			product.TradeId = trade.TradeId
			product.StockAmount.HasValue = false // set the stock to inf
			trades[i].Products.Slice[j] = product
		}
		_, err = session.Table("s_trade_product").Insert(trades[i].Products.Slice)
		utils2.CheckErr(err)
	}
	_, err = session.Table("s_trade").Insert(trades)
	utils2.CheckErr(err)
}

func init() {
	addTable("s_trade", client.Trade{}, InsertTrade)
	addTable("s_trade_product", client.TradeProduct{}, nil)
}
