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

	for id, card := range gamedata.Instance.Card {
		userCard := user_card.GetUserCard(session, id)
		entry := response.WebUICardEntry{
			Id:         id,
			Name:       dictionary.Resolve(card.Appearance.CardName),
			MemberId:   *card.MemberMasterId,
			MemberName: dictionary.Resolve(card.Member.Name),
			GroupId:    card.Member.MemberGroupId,
			GroupName:  dictionary.Resolve(card.Member.MemberGroup.GroupName),
			Rarity:     card.Rarity.CardRarityType,
			Grade:      userCard.Grade,
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
