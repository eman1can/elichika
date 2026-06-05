package user

import (
	"net/http"
	"sort"

	"elichika/internal/subsystem/user_voice"

	"elichika/internal/gamedata"
	"elichika/internal/server"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

type WebUIListVoiceRequest struct {
	ReleaseRoute *int32 `form:"route"`
	ListType     *int32 `form:"list_type"`
}

type WebUINaviVoiceEntry struct {
	Id             int32  `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	SoundAssetPath string `json:"sound_asset_path"`
	MemberId       int32  `json:"member_id"`
	MemberName     string `json:"member_name"`
	GroupId        int32  `json:"group_id"`
	GroupName      string `json:"group_name"`
	DisplayOrder   int32  `json:"display_order"`
	ReleaseRoute   int32  `json:"release_route"`
	ReleaseValue   int32  `json:"release_value"`
	Owned          bool   `json:"owned"`
}

func naviVoiceList(ctx *gin.Context) {
	var req WebUIListVoiceRequest
	var resp []WebUINaviVoiceEntry

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := ctx.MustGet("session").(*userdata.Session)
	dictionary := ctx.MustGet("dictionary").(*gamedata.Dictionary)

	route := req.ReleaseRoute
	listType := req.ListType

	for id, voice := range session.Gamedata.NaviVoice {
		userVoice := user_voice.GetUserVoice(session, id)

		member := session.Gamedata.Member[voice.MemberMId]

		if route != nil && voice.NaviVoiceReleaseRoute != *route {
			continue
		}

		if listType != nil && voice.ListType == *listType {
			continue
		}

		entry := WebUINaviVoiceEntry{
			Id:             id,
			Name:           dictionary.Resolve(voice.Name),
			Description:    dictionary.Resolve(voice.Description),
			SoundAssetPath: voice.SheetName,
			MemberId:       voice.MemberMId,
			MemberName:     dictionary.Resolve(member.Name),
			GroupId:        member.MemberGroupId,
			GroupName:      dictionary.Resolve(member.MemberGroup.GroupName),
			DisplayOrder:   voice.DisplayOrder,
			ReleaseRoute:   voice.NaviVoiceReleaseRoute,
			ReleaseValue:   voice.NaviVoiceReleaseValue,
			Owned:          !userVoice.IsNew,
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

	ctx.JSON(http.StatusOK, resp)
}

func init() {
	server.AddHandler("/webui/user", "GET", "/voice", naviVoiceList)
}
