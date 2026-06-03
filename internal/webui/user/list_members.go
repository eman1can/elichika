package user

import (
	"encoding/json"
	"net/http"
	"sort"

	"elichika/internal/gamedata"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_member"
	"elichika/internal/subsystem/user_story_member"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

type WebUIListMembersRequest struct {
	Language string `form:"l" json:"l"`
}

type WebUIMemberEntry struct {
	Id                 int32  `json:"member_id"`
	Name               string `json:"member_name"`
	GroupId            int32  `json:"group_id"`
	GroupName          string `json:"group_name"`
	ImageAssetPath     string `json:"image_asset_path"`
	LoveLevel          int32  `json:"love_level"`
	IsLoveLevelMaxed   bool   `json:"is_love_level_maxed"`
	IsLovePanelMaxed   bool   `json:"is_love_panel_maxed"`
	IsAllStoryUnlocked bool   `json:"is_all_story_unlocked"`
}

func memberList(ctx *gin.Context) {
	var req WebUIListMembersRequest
	var resp []WebUIMemberEntry

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := ctx.MustGet("session").(*userdata.Session)
	dictionary := gamedata.DictionaryByLanguage(req.Language)

	for id, member := range gamedata.Instance.Member {
		userMember := user_member.GetMember(session, id)
		entry := WebUIMemberEntry{
			Id:                 id,
			Name:               dictionary.Resolve(member.Name),
			GroupId:            member.MemberGroupId,
			GroupName:          dictionary.Resolve(member.MemberGroup.GroupName),
			LoveLevel:          userMember.LoveLevel,
			IsLoveLevelMaxed:   userMember.LoveLevel >= session.Gamedata.MemberLoveLevelMax || userMember.LovePoint >= userMember.LovePointLimit,
			IsLovePanelMaxed:   userMember.IsCurrentLovePanelMaxed || userMember.IsAllLovePanelMaxed,
			IsAllStoryUnlocked: user_story_member.AllStoryMembersUnlocked(session, id),
		}

		if cards, ok := gamedata.Instance.CardByMemberId[id]; ok && len(cards) > 0 {
			entry.ImageAssetPath = cards[0].IdolAppearance.ThumbnailAssetPath
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
