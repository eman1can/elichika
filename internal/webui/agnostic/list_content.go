package agnostic

import (
	"elichika/internal/enum"
	"elichika/internal/server"
	"elichika/internal/utils"
	"elichika/internal/webui/response"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	GenericContentTypeToName = map[int32]string{}
)

func listContent(ctx *gin.Context) {
	resp := response.WebUIContentListResponse{}

	for contentType, name := range GenericContentTypeToName {
		resp.Items = append(resp.Items, response.WebUIContent{
			ContentType: contentType,
			Name:        name,
		})
	}

	jsonBytes, err := json.Marshal(resp)
	utils.CheckErr(err)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, string(jsonBytes))
}

func init() {
	// TODO: Add these items to the dictionary and Masterdata?
	GenericContentTypeToName[enum.ContentTypeSnsCoin] = "Gold"
	GenericContentTypeToName[enum.ContentTypeCard] = "Member Card"
	GenericContentTypeToName[enum.ContentTypeCardExp] = "Member Card EXP"
	GenericContentTypeToName[enum.ContentTypeGachaPoint] = "Shiny Quartz"
	GenericContentTypeToName[enum.ContentTypeLessonEnhancingItem] = "Lesson Enhancing Item"
	GenericContentTypeToName[enum.ContentTypeSuit] = "Member Suit"
	GenericContentTypeToName[enum.ContentTypeVoice] = "Member Voice"
	GenericContentTypeToName[enum.ContentTypeGachaTicket] = "Gacha Ticket"
	GenericContentTypeToName[enum.ContentTypeGameMoney] = "Money"
	GenericContentTypeToName[enum.ContentTypeTrainingMaterial] = "Training Material"
	GenericContentTypeToName[enum.ContentTypeGradeUpper] = "Rarity Increase Item"
	GenericContentTypeToName[enum.ContentTypeGiftBox] = "Gift Box"
	GenericContentTypeToName[enum.ContentTypeEmblem] = "Member Emblem"
	GenericContentTypeToName[enum.ContentTypeRecoveryAp] = "Training Ticket"
	GenericContentTypeToName[enum.ContentTypeRecoveryLp] = "Show Candy"
	GenericContentTypeToName[enum.ContentTypeStorySide] = "Side Story"
	GenericContentTypeToName[enum.ContentTypeStoryMember] = "Side Story Member"
	GenericContentTypeToName[enum.ContentTypeExchangeEventPoint] = "Exchange Event Point"
	GenericContentTypeToName[enum.ContentTypeAccessory] = "Accessory"
	GenericContentTypeToName[enum.ContentTypeAccessoryLevelUp] = "Accessory Level Up"
	GenericContentTypeToName[enum.ContentTypeAccessoryRarityUp] = "Accessory Rarity Up"
	GenericContentTypeToName[enum.ContentTypeCustomBackground] = "Background"
	GenericContentTypeToName[enum.ContentTypeEventMarathonBooster] = "Event Marathon Booster"
	GenericContentTypeToName[enum.ContentTypeLiveSkipTicket] = "Live Skip Ticket"
	GenericContentTypeToName[enum.ContentTypeEventMiningBooster] = "Event Mining Booster"
	GenericContentTypeToName[enum.ContentTypeStoryEventUnlock] = "Story Event Unlock Key"
	GenericContentTypeToName[enum.ContentTypeRecoveryTowerCardUsedCount] = "Recovery Tower Card Used Count"
	GenericContentTypeToName[enum.ContentTypeSubscriptionCoin] = "Subscription Coin"
	GenericContentTypeToName[enum.ContentTypeMemberGuildSupport] = "Member Guild Support"

	server.AddHandler("/webui", "GET", "/content", listContent)
}
