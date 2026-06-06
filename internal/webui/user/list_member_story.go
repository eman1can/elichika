package user

import (
	"net/http"

	"elichika/internal/gamedata"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_member"
	"elichika/internal/subsystem/user_story_member"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

func listMemberStory(ctx *gin.Context) {
	var resp []WebUIStoryChapterEntry

	session := ctx.MustGet("session").(*userdata.Session)
	dictionary := ctx.MustGet("dictionary").(*gamedata.Dictionary)

	for memberId, masterMember := range session.Gamedata.Member {
		entry := WebUIStoryChapterEntry{
			Id:           memberId,
			Title:        dictionary.Resolve(masterMember.Name),
			Description:  dictionary.Resolve(masterMember.MemberGroup.GroupName),
			DisplayOrder: memberId,
			Chapters:     make([]WebUIStoryCellEntry, 0),
		}

		if cards, ok := session.Gamedata.CardByMemberId[memberId]; ok && len(cards) > 0 {
			entry.ImageAssetPath = cards[0].IdolAppearance.ThumbnailAssetPath
		}

		member := user_member.GetMember(session, memberId)

		for _, storyMember := range session.Gamedata.StoryMember {
			if storyMember.MemberMId != memberId {
				continue
			}

			entry.Chapters = append(entry.Chapters, WebUIStoryCellEntry{
				Id:             storyMember.Id,
				Chapter:        storyMember.StoryNumber,
				Title:          dictionary.Resolve(storyMember.Title),
				Description:    dictionary.Resolve(storyMember.Description),
				ImageAssetPath: storyMember.CardImageAssetPath,
				IsNew:          user_story_member.GetStoryMember(session, storyMember.Id).IsNew,
				Unlocked:       member.LoveLevel >= storyMember.LoveLevel,
			})
		}

		resp = append(resp, entry)
	}

	ctx.JSON(http.StatusOK, resp)
}

func init() {
	server.AddHandler("/webui/user", "GET", "/list_member_story", listMemberStory)
}
