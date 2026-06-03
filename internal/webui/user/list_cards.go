package user

import (
	"encoding/json"
	"net/http"
	"sort"

	"elichika/internal/gamedata"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_card"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

type WebUIListCardsRequest struct {
	Language string `form:"l" json:"l"`
}

type WebUICardEntry struct {
	Id                     int32  `json:"card_id"`
	ImageAssetPath         string `json:"image_asset_path"`
	Name                   string `json:"card_name"`
	MemberId               int32  `json:"member_id"`
	MemberName             string `json:"member_name"`
	GroupId                int32  `json:"group_id"`
	GroupName              string `json:"group_name"`
	Grade                  int32  `json:"grade"`
	Rarity                 int32  `json:"rarity"`
	IsAllTrainingActivated bool   `json:"is_all_training_activated"`
}

func cardList(ctx *gin.Context) {
	var req WebUIListCardsRequest
	var resp []WebUICardEntry

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := ctx.MustGet("session").(*userdata.Session)
	dictionary := gamedata.DictionaryByLanguage(req.Language)

	for id, masterCard := range gamedata.Instance.Card {
		card := user_card.GetUserCard(session, id)
		entry := WebUICardEntry{
			Id:                     id,
			Name:                   dictionary.Resolve(masterCard.Appearance.CardName),
			ImageAssetPath:         masterCard.IdolAppearance.ThumbnailAssetPath,
			MemberId:               *masterCard.MemberMasterId,
			MemberName:             dictionary.Resolve(masterCard.Member.Name),
			GroupId:                masterCard.Member.MemberGroupId,
			GroupName:              dictionary.Resolve(masterCard.Member.MemberGroup.GroupName),
			Rarity:                 masterCard.Rarity.CardRarityType,
			Grade:                  card.Grade,
			IsAllTrainingActivated: card.IsAllTrainingActivated,
		}
		resp = append(resp, entry)
	}

	sort.Slice(resp, func(i, j int) bool {
		if resp[i].GroupId != resp[j].GroupId {
			return resp[i].GroupId < resp[j].GroupId
		}
		return resp[i].Id < resp[j].Id
	})

	jsonBytes, err := json.Marshal(resp)
	utils.CheckErr(err)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, string(jsonBytes))
}

func init() {
	server.AddHandler("/webui/user", "GET", "/card", cardList)
}
