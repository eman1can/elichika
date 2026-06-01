package response

import (
	"elichika/internal/client"
)

type WebUIUserAccountInfo struct {
	UserId                int32                `xorm:"pk 'user_id'" json:"user_id"`
	Name                  client.LocalizedText `xorm:"'name'" json:"name"`
	Nickname              client.LocalizedText `xorm:"'nickname'" json:"nickname"`
	LastLoginAt           int64                `json:"last_login_at"`
	Rank                  int32                `json:"rank"`
	Exp                   int32                `json:"exp"`
	Message               client.LocalizedText `xorm:"'message'" json:"message"`
	RecommendCardMasterId int32                `xorm:"'recommend_card_master_id'" json:"recommend_card_master_id"`
}

type WebUIAccountListResponse struct {
	Accounts []WebUIUserAccountInfo `json:"accounts"`
}
