package gamedata

import (
	"log"

	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/utils"

	"xorm.io/xorm"
)

type Content struct {
	Content   client.Content
	Name      string `json:"name"`
	AssetPath string `json:"asset_path"`
}

func loadContentFromTable(gamedata *Gamedata, contentType int32, table string) {
	type genericContent struct {
		Id        int32  `xorm:"pk 'id'"`
		Name      string `xorm:"name"`
		AssetPath string `xorm:"thumbnail_asset_path"`
	}
	var contents []genericContent

	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table(table).Find(&contents)
	})
	utils.CheckErr(err)

	for _, content := range contents {
		gamedata.Content[contentType][content.Id] = &Content{
			Content: client.Content{
				ContentType: contentType,
				ContentId:   content.Id,
			},
			Name:      content.Name,
			AssetPath: content.AssetPath,
		}
	}
}

func loadContentFromUITextureTable(gamedata *Gamedata, contentType int32, textureKey int32, name string) {
	var assetPath string
	var exists bool
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		exists, err = session.Table("m_ui_texture").Where("id = ?", textureKey).Cols("asset_path").Get(&assetPath)
	})
	utils.CheckErrMustExist(err, exists)

	gamedata.Content[contentType][0] = &Content{
		Content: client.Content{
			ContentType: contentType,
			ContentId:   0,
		},
		Name:      name,
		AssetPath: assetPath,
	}
}

func loadContent(gamedata *Gamedata) {
	log.Println("Loading Content")
	gamedata.Content = make(map[int32]map[int32]*Content)
	for contentType := range gamedata.ContentType {
		gamedata.Content[contentType] = make(map[int32]*Content)
	}

	loadContentFromTable(gamedata, enum.ContentTypeGachaPoint, "m_gacha_point_setting")
	loadContentFromTable(gamedata, enum.ContentTypeLessonEnhancingItem, "m_lesson_enhancing_item")
	loadContentFromTable(gamedata, enum.ContentTypeGachaTicket, "m_gacha_ticket")
	loadContentFromTable(gamedata, enum.ContentTypeTrainingMaterial, "m_training_material")
	loadContentFromTable(gamedata, enum.ContentTypeGradeUpper, "m_grade_upper")
	loadContentFromTable(gamedata, enum.ContentTypeRecoveryAp, "m_recovery_ap")
	loadContentFromTable(gamedata, enum.ContentTypeRecoveryLp, "m_recovery_lp")
	loadContentFromTable(gamedata, enum.ContentTypeExchangeEventPoint, "m_exchange_event_point")
	loadContentFromTable(gamedata, enum.ContentTypeAccessoryLevelUp, "m_accessory_level_up_item")
	loadContentFromTable(gamedata, enum.ContentTypeAccessoryRarityUp, "m_accessory_rarity_up_item")
	loadContentFromTable(gamedata, enum.ContentTypeEventMarathonBooster, "m_event_marathon_booster_item")
	loadContentFromTable(gamedata, enum.ContentTypeLiveSkipTicket, "m_live_skip_ticket")
	loadContentFromTable(gamedata, enum.ContentTypeStoryEventUnlock, "m_story_event_unlock_item")
	loadContentFromTable(gamedata, enum.ContentTypeRecoveryTowerCardUsedCount, "m_recovery_tower_card_used_count_item")

	loadContentFromUITextureTable(gamedata, enum.ContentTypeSnsCoin, enum.UiTextureKeySnsCoinIcon, "k.item_name_21059")
	loadContentFromUITextureTable(gamedata, enum.ContentTypeCardExp, enum.UiTextureKeyCardExpIcon, "k.item_card_exp_name")
	loadContentFromUITextureTable(gamedata, enum.ContentTypeGameMoney, enum.UiTextureKeyGameMoneyIcon, "k.item_name_1201")
	loadContentFromUITextureTable(gamedata, enum.ContentTypeVoice, enum.UiTextureKeyVoiceIcon, "Member Voice")
	loadContentFromUITextureTable(gamedata, enum.ContentTypeStorySide, enum.UiTextureKeyStoryMember, "Side Story Episode")
	loadContentFromUITextureTable(gamedata, enum.ContentTypeStoryMember, enum.UiTextureKeyStoryMember, "Member Side Story Episode")
	loadContentFromUITextureTable(gamedata, enum.ContentTypeSubscriptionCoin, enum.UiTextureKeySubscriptionCoinIcon, "Subscription Coin")

	// TODO: Ensure that cards / suits / emblems load data into the ContentsByContentType table!!!

	gamedata.ContentsByContentType = make(map[int32][]*client.Content)
	for contentType := range gamedata.Content {
		for _, content := range gamedata.Content[contentType] {
			gamedata.ContentsByContentType[contentType] = append(gamedata.ContentsByContentType[contentType], &content.Content)
		}
	}
}

func init() {
	addLoadFunc(loadContent)
	addPrequisite(loadContent, loadContentType)
}
