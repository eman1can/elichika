package user

import (
	"encoding/json"
	"net/http"
	"sort"

	"elichika/internal/subsystem/user_voice"

	"elichika/internal/gamedata"
	"elichika/internal/server"
	"elichika/internal/userdata"
	"elichika/internal/utils"
	"elichika/internal/webui/request"
	"elichika/internal/webui/response"

	"github.com/gin-gonic/gin"
)

func naviVoiceList(ctx *gin.Context) {
	resp := response.WebUINaviVoiceListResponse{}
	req := request.WebUIListVoiceRequest{}
	err := ctx.ShouldBindQuery(&req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)
	dictionary := gamedata.DictionaryByLanguage(req.Language)

	route := req.ReleaseRoute
	listType := req.ListType

	for id, voice := range gamedata.Instance.NaviVoice {
		userVoice := user_voice.GetUserVoice(session, id)

		member := session.Gamedata.Member[voice.MemberMId]

		if route != nil && voice.NaviVoiceReleaseRoute != *route {
			continue
		}

		if listType != nil && voice.ListType == *listType {
			continue
		}

		entry := response.WebUINaviVoiceEntry{
			Id:           id,
			Name:         dictionary.Resolve(voice.Name),
			Description:  dictionary.Resolve(voice.Description),
			MemberId:     voice.MemberMId,
			MemberName:   dictionary.Resolve(member.Name),
			GroupId:      member.MemberGroupId,
			GroupName:    dictionary.Resolve(member.MemberGroup.GroupName),
			DisplayOrder: voice.DisplayOrder,
			ReleaseRoute: voice.NaviVoiceReleaseRoute,
			ReleaseValue: voice.NaviVoiceReleaseValue,
			Owned:        !userVoice.IsNew,
		}
		resp = append(resp, entry)
	}

	sort.Slice(resp, func(i, j int) bool {
		if resp[i].GroupId != resp[j].GroupId {
			return resp[i].GroupId < resp[j].GroupId
		}
		if resp[i].MemberId != resp[j].MemberId {
			return resp[i].MemberId < resp[j].MemberId
		}
		return resp[i].DisplayOrder < resp[j].DisplayOrder
	})

	jsonBytes, err := json.Marshal(resp)
	utils.CheckErr(err)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, string(jsonBytes))
}

func init() {
	server.AddHandler("/webui/user", "GET", "/voice", naviVoiceList)
}
