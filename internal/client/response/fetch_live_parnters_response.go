package response

import "elichika/internal/client"

type FetchLiveParntersResponse struct {
	PartnerSelectState client.PartnerSelectState `json:"partner_select_state"`
}
