package user

import (
	"net/http"
	"sort"

	"elichika/internal/gamedata"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_suit"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

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
	var resp []WebUISuitEntry

	session := ctx.MustGet("session").(*userdata.Session)
	dictionary := ctx.MustGet("dictionary").(*gamedata.Dictionary)

	for id, suit := range session.Gamedata.Suit {
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

	ctx.JSON(http.StatusOK, resp)
}

func init() {
	server.AddHandler("/webui/user", "GET", "/suit", suitList)
}
