package user

import (
	"encoding/json"
	"net/http"
	"sort"

	"elichika/internal/gamedata"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_suit"
	"elichika/internal/userdata"
	"elichika/internal/utils"
	"elichika/internal/webui/request"
	"elichika/internal/webui/response"

	"github.com/gin-gonic/gin"
)

func suitList(ctx *gin.Context) {
	resp := response.WebUISuitListResponse{}
	req := request.WebUILanguageRequest{}
	err := ctx.ShouldBindQuery(&req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)
	dictionary := gamedata.DictionaryByLanguage(req.Language)

	for id, suit := range gamedata.Instance.Suit {
		entry := response.WebUISuitEntry{
			Id:         id,
			Name:       dictionary.Resolve(suit.Name),
			MemberId:   *suit.MemberMId,
			MemberName: dictionary.Resolve(suit.Member.Name),
			GroupId:    suit.Member.MemberGroupId,
			GroupName:  dictionary.Resolve(suit.Member.MemberGroup.GroupName),
			Owned:      user_suit.HasSuit(session, id),
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
	server.AddHandler("/webui/user", "GET", "/suit", suitList)
}
