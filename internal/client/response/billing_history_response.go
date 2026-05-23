package response

import (
	"elichika/internal/client"
	"elichika/internal/generic"
)

type BillingHistoryResponse struct {
	BillingBalanceHistoryList generic.List[client.BillingBalanceHistory] `json:"billing_balance_history_list"`
	BillingDepositHistoryList generic.List[client.BillingDepositHistory] `json:"billing_deposit_history_list"`
	BillingConsumeHistoryList generic.List[client.BillingConsumeHistory] `json:"billing_consume_history_list"`
}
