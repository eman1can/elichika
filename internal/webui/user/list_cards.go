package user

import (
	"net/http"
	"sort"

	"elichika/internal/gamedata"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_card"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

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
	var resp []WebUICardEntry

	session := ctx.MustGet("session").(*userdata.Session)
	dictionary := ctx.MustGet("dictionary").(*gamedata.Dictionary)

	for id, masterCard := range session.Gamedata.Card {
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

	ctx.JSON(http.StatusOK, resp)
}

func init() {
	server.AddHandler("/webui/user", "GET", "/card", cardList)
}
