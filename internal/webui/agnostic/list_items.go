package agnostic

import (
	"log"
	"net/http"

	"elichika/internal/enum"
	"elichika/internal/gamedata"
	"elichika/internal/server"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

type Item struct {
	Id        int32  `json:"id"`
	Name      string `json:"name"`
	AssetPath string `json:"asset_path"`
}

var (
	ItemsByItemId = map[int32]map[int32]Item{}
)

type WebUIItem struct {
	Name        string `json:"name"`
	ContentType int32  `json:"content_type"`
	ContentId   int32  `json:"content_id"`
}

func loadItemsFromTableWithName(contentType int32, tableName string) {
	log.Println("Trying to load items from table", tableName)
	type ContentTypeItem struct {
		Id                 int32  `xorm:"id"`
		Name               string `xorm:"name"`
		ThumbnailAssetPath string `xorm:"thumbnail_asset_path"`
	}

	var contents []ContentTypeItem
	gamedata.Instance.MasterdataDb.Do(func(session *xorm.Session) {
		err := session.Table(tableName).Find(&contents)
		utils.CheckErr(err)
	})

	assets := map[int32]Item{}
	for _, item := range contents {
		assets[item.Id] = Item{
			Id:        item.Id,
			Name:      item.Name,
			AssetPath: item.ThumbnailAssetPath,
		}
	}
	ItemsByItemId[contentType] = assets
}

func loadItemsFromTable(contentType int32, tableName string) {
	log.Println("Trying to load items from table", tableName)
	type ContentTypeItem struct {
		Id                 int32  `xorm:"id"`
		ThumbnailAssetPath string `xorm:"thumbnail_asset_path"`
	}

	var contents []ContentTypeItem
	gamedata.Instance.MasterdataDb.Do(func(session *xorm.Session) {
		err := session.Table(tableName).Find(&contents)
		utils.CheckErr(err)
	})

	assets := map[int32]Item{}
	for _, item := range contents {

		assets[item.Id] = Item{
			Id:        item.Id,
			Name:      GenericContentTypeToName[contentType],
			AssetPath: item.ThumbnailAssetPath,
		}
	}
	ItemsByItemId[contentType] = assets
}

func loadItemFromUITextureTable(contentType int32, textureKey int32) {
	var assetPath string
	gamedata.Instance.MasterdataDb.Do(func(session *xorm.Session) {
		exists, err := session.Table("m_ui_texture").Where("id = ?", textureKey).Cols("asset_path").Get(&assetPath)
		utils.CheckErrMustExist(err, exists)
	})

	ItemsByItemId[contentType] = map[int32]Item{0: Item{
		Id:        0,
		Name:      GenericContentTypeToName[contentType],
		AssetPath: assetPath,
	}}
}

type WebUIItemListRequest struct {
	ContentType int32 `form:"type"`
}

func listItems(ctx *gin.Context) {
	var req WebUIItemListRequest
	var resp []WebUIItem

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dictionary := ctx.MustGet("dictionary").(*gamedata.Dictionary)

	for itemId, item := range ItemsByItemId[req.ContentType] {
		resp = append(resp, WebUIItem{
			ContentType: req.ContentType,
			Name:        dictionary.Resolve(item.Name),
			ContentId:   itemId,
		})
	}

	ctx.JSON(http.StatusOK, resp)
}

func init() {
	var genericContentTypeToTableName = map[int32]string{}
	genericContentTypeToTableName[enum.ContentTypeGachaPoint] = "m_gacha_point"
	genericContentTypeToTableName[enum.ContentTypeLessonEnhancingItem] = "m_lesson_enhancing_item"
	genericContentTypeToTableName[enum.ContentTypeGachaTicket] = "m_gacha_ticket"
	genericContentTypeToTableName[enum.ContentTypeTrainingMaterial] = "m_training_material"
	genericContentTypeToTableName[enum.ContentTypeGradeUpper] = "m_grade_upper"
	genericContentTypeToTableName[enum.ContentTypeRecoveryAp] = "m_recovery_ap"
	genericContentTypeToTableName[enum.ContentTypeRecoveryLp] = "m_recovery_lp"
	genericContentTypeToTableName[enum.ContentTypeExchangeEventPoint] = "m_exchange_event_point"
	genericContentTypeToTableName[enum.ContentTypeAccessory] = "m_accessory"
	genericContentTypeToTableName[enum.ContentTypeAccessoryLevelUp] = "m_accessory_level_up_item"
	genericContentTypeToTableName[enum.ContentTypeAccessoryRarityUp] = "m_accessory_rarity_up_item"
	genericContentTypeToTableName[enum.ContentTypeCustomBackground] = "m_custom_background"
	genericContentTypeToTableName[enum.ContentTypeEventMarathonBooster] = "m_event_marathon_booster_item"
	genericContentTypeToTableName[enum.ContentTypeLiveSkipTicket] = "m_live_skip_ticket"
	genericContentTypeToTableName[enum.ContentTypeEventMiningBooster] = "m_event_mining_booster_item"
	genericContentTypeToTableName[enum.ContentTypeStoryEventUnlock] = "m_story_event_unlock_item"
	genericContentTypeToTableName[enum.ContentTypeRecoveryTowerCardUsedCount] = "m_recovery_tower_card_used_count_item"
	genericContentTypeToTableName[enum.ContentTypeMemberGuildSupport] = "m_member_guild_support_item"

	var contentTypeToUiTextureKey = map[int32]int32{}
	contentTypeToUiTextureKey[enum.ContentTypeSnsCoin] = enum.UiTextureKeySnsCoinIcon
	contentTypeToUiTextureKey[enum.ContentTypeCardExp] = enum.UiTextureKeyCardExpIcon
	contentTypeToUiTextureKey[enum.ContentTypeGameMoney] = enum.UiTextureKeyGameMoneyIcon
	contentTypeToUiTextureKey[enum.ContentTypeVoice] = enum.UiTextureKeyVoiceIcon
	contentTypeToUiTextureKey[enum.ContentTypeStorySide] = enum.UiTextureKeyStoryMember
	contentTypeToUiTextureKey[enum.ContentTypeStoryMember] = enum.UiTextureKeyStoryMember
	contentTypeToUiTextureKey[enum.ContentTypeSubscriptionCoin] = enum.UiTextureKeySubscriptionCoinIcon

	for contentType := range gamedata.Instance.ContentType {
		switch contentType {
		case enum.ContentTypeLessonEnhancingItem:
			fallthrough
		case enum.ContentTypeGachaTicket:
			fallthrough
		case enum.ContentTypeTrainingMaterial:
			fallthrough
		case enum.ContentTypeGradeUpper:
			fallthrough
		case enum.ContentTypeRecoveryAp:
			fallthrough
		case enum.ContentTypeRecoveryLp:
			fallthrough
		case enum.ContentTypeExchangeEventPoint:
			fallthrough
		case enum.ContentTypeAccessory:
			fallthrough
		case enum.ContentTypeAccessoryLevelUp:
			fallthrough
		case enum.ContentTypeAccessoryRarityUp:
			fallthrough
		case enum.ContentTypeCustomBackground:
			fallthrough
		case enum.ContentTypeEventMarathonBooster:
			fallthrough
		case enum.ContentTypeLiveSkipTicket:
			fallthrough
		case enum.ContentTypeEventMiningBooster:
			fallthrough
		case enum.ContentTypeStoryEventUnlock:
			fallthrough
		case enum.ContentTypeRecoveryTowerCardUsedCount:
			fallthrough
		case enum.ContentTypeMemberGuildSupport:
			tableName := genericContentTypeToTableName[contentType]
			loadItemsFromTableWithName(contentType, tableName)
		case enum.ContentTypeGachaPoint:
			tableName := genericContentTypeToTableName[contentType]
			loadItemsFromTable(contentType, tableName)
		case enum.ContentTypeSnsCoin:
			fallthrough
		case enum.ContentTypeCardExp:
			fallthrough
		case enum.ContentTypeGameMoney:
			fallthrough
		case enum.ContentTypeVoice:
			fallthrough
		case enum.ContentTypeStorySide:
			fallthrough
		case enum.ContentTypeStoryMember:
			fallthrough
		case enum.ContentTypeSubscriptionCoin:
			textureKey := contentTypeToUiTextureKey[contentType]
			loadItemFromUITextureTable(contentType, textureKey)
		case enum.ContentTypeSuit:
			assets := map[int32]Item{}
			for suitMasterId, suit := range gamedata.Instance.Suit {
				assets[suitMasterId] = Item{
					Id:        suitMasterId,
					Name:      suit.Name,
					AssetPath: suit.ThumbnailImageAssetPath,
				}
			}
			ItemsByItemId[contentType] = assets
		case enum.ContentTypeCard:
			assets := map[int32]Item{}
			for cardMasterId, card := range gamedata.Instance.Card {
				assets[cardMasterId] = Item{
					Id:        cardMasterId,
					Name:      card.IdolAppearance.CardName,
					AssetPath: card.IdolAppearance.ThumbnailAssetPath,
				}
			}
			ItemsByItemId[contentType] = assets
		case enum.ContentTypeEmblem:
			assets := map[int32]Item{}
			for emblemMasterId, emblem := range gamedata.Instance.Emblem {
				assets[emblemMasterId] = Item{
					Id:        emblemMasterId,
					Name:      emblem.Name,
					AssetPath: emblem.EmblemAssetPath,
				}
			}
			ItemsByItemId[contentType] = assets
		}
	}

	server.AddHandler("/webui", "GET", "/item", listItems)
}
