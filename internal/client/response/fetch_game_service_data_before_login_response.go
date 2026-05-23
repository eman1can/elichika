package response

import (
	"elichika/internal/client"
	"elichika/internal/generic"
)

type FetchGameServiceDataBeforeLoginResponse struct {
	Data generic.Nullable[client.UserLinkDataBeforeLogin] `json:"data"`
}
