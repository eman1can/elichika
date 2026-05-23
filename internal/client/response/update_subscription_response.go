package response

import "elichika/internal/client"

type UpdateSubscriptionResponse struct {
	UserModel        *client.UserModel       `json:"user_model"`
	BillingStateInfo client.BillingStateInfo `json:"billing_state_info"`
}
