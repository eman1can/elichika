package user_tower

import (
	"elichika/internal/client"
	"elichika/internal/userdata"
	"elichika/internal/utils"
)

func GetUserTowerCardUsed(session *userdata.Session, towerId, cardMasterId int32) client.TowerCardUsedCount {
	cardUsed := client.TowerCardUsedCount{}
	exist, err := session.Db.Table("u_tower_card_used_count").
		Where("user_id = ? AND tower_id = ? AND card_master_id = ?", session.UserId, towerId, cardMasterId).Get(&cardUsed)
	utils.CheckErr(err)
	if !exist {
		cardUsed = client.TowerCardUsedCount{
			CardMasterId:   cardMasterId,
			UsedCount:      0,
			RecoveredCount: 0,
			LastUsedAt:     0,
		}
	}
	return cardUsed
}
