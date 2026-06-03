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

	"github.com/gin-gonic/gin"
)

type WebUIListSuitsRequest struct {
	Language string `form:"l" json:"l"`
}

type WebUISuitEntry struct {
	Id             int32  `json:"suit_id"`
	Name           string `json:"suit_name"`
	ImageAssetPath string `json:"image_asset_path"`
	MemberId       int32  `json:"member_id"`
	MemberName     string `json:"member_name"`
	GroupId        int32  `json:"group_id"`
	GroupName      string `json:"group_name"`
	Owned          bool   `json:"owned"`
}

func suitList(ctx *gin.Context) {
	var req WebUIListSuitsRequest
	var resp []WebUISuitEntry

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := ctx.MustGet("session").(*userdata.Session)
	dictionary := gamedata.DictionaryByLanguage(req.Language)

	for id, suit := range gamedata.Instance.Suit {
		entry := WebUISuitEntry{
			Id:             id,
			Name:           dictionary.Resolve(suit.Name),
			ImageAssetPath: suit.ThumbnailImageAssetPath,
			MemberId:       *suit.MemberMId,
			MemberName:     dictionary.Resolve(suit.Member.Name),
			GroupId:        suit.Member.MemberGroupId,
			GroupName:      dictionary.Resolve(suit.Member.MemberGroup.GroupName),
			Owned:          user_suit.HasSuit(session, id),
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
