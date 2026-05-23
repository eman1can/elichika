package response

import (
	"elichika/internal/client"
	"elichika/internal/generic"
)

type GetVoltageRankingResponse struct {
	VoltageRankingCells generic.List[client.VoltageRankingCell] `json:"voltage_ranking_cells"`
}
