package user

import (
	"encoding/json"
	"net/http"
	"sort"

	"elichika/internal/gamedata"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_member"
	"elichika/internal/userdata"
	"elichika/internal/utils"
	"elichika/internal/webui/request"
	"elichika/internal/webui/response"

	"github.com/gin-gonic/gin"
)

func memberList(ctx *gin.Context) {
	resp := response.WebUIMemberListResponse{}
	req := request.WebUILanguageRequest{}
	err := ctx.ShouldBindQuery(&req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)
	dictionary := gamedata.DictionaryByLanguage(req.Language)

	for id, member := range gamedata.Instance.Member {
		userMember := user_member.GetMember(session, id)
		entry := response.WebUIMemberEntry{
			Id:        id,
			Name:      dictionary.Resolve(member.Name),
			GroupId:   member.MemberGroupId,
			GroupName: dictionary.Resolve(member.MemberGroup.GroupName),
			LoveLevel: userMember.LoveLevel,
		}
		if cards, ok := gamedata.Instance.CardByMemberId[id]; ok && len(cards) > 0 {
			entry.RepresentativeCardId = cards[0].Id
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
	server.AddHandler("/webui/user", "GET", "/member", memberList)
}
