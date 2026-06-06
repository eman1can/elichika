package user

import (
	"net/http"
	"sort"

	"elichika/internal/gamedata"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_custom_background"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

type WebUIBackgroundEntry struct {
	Id             int32  `json:"id"`
	Name           string `json:"name"`
	ImageAssetPath string `json:"image_asset_path"`
	DisplayOrder   int32  `json:"display_order"`
	Owned          bool   `json:"owned"`
}

func backgroundList(ctx *gin.Context) {
	var resp []WebUIBackgroundEntry

	dictionary := ctx.MustGet("dictionary").(*gamedata.Dictionary)
	session := ctx.MustGet("session").(*userdata.Session)

	for id, background := range session.Gamedata.CustomBackground {
		userBackground := user_custom_background.GetUserCustomBackground(session, id)
		entry := WebUIBackgroundEntry{
			Id:             id,
			Name:           dictionary.Resolve(background.Name),
			ImageAssetPath: background.ThumbnailAssetPath,
			DisplayOrder:   background.DisplayOrder,
			Owned:          !userBackground.IsNew,
		}
		resp = append(resp, entry)
	}

	sort.Slice(resp, func(i, j int) bool {
		if resp[i].DisplayOrder != resp[j].DisplayOrder {
			return resp[i].DisplayOrder < resp[j].DisplayOrder
		}
		return resp[i].Id < resp[j].Id
	})

	ctx.JSON(http.StatusOK, resp)
}

func init() {
	server.AddHandler("/webui/user", "GET", "/background", backgroundList)
}
