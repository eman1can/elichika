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
	"elichika/internal/webui/request"
	"elichika/internal/webui/response"

	"github.com/gin-gonic/gin"
)

func cardList(ctx *gin.Context) {
	resp := response.WebUICardListResponse{}
	req := request.WebUILanguageRequest{}
	err := ctx.ShouldBindQuery(&req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)
	dictionary := gamedata.DictionaryByLanguage(req.Language)

	for id, masterCard := range gamedata.Instance.Card {
		card := user_card.GetUserCard(session, id)
		entry := response.WebUICardEntry{
			Id:                     id,
			Name:                   dictionary.Resolve(masterCard.Appearance.CardName),
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
