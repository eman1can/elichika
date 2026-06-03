package user

import (
	"encoding/json"
	"net/http"
	"sort"

	"elichika/internal/gamedata"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_custom_background"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

type WebUIListBackgroundRequest struct {
	Language string `form:"l" json:"l"`
}

type WebUIBackgroundEntry struct {
	Id             int32  `json:"id"`
	Name           string `json:"name"`
	ImageAssetPath string `json:"image_asset_path"`
	DisplayOrder   int32  `json:"display_order"`
	Owned          bool   `json:"owned"`
}

func backgroundList(ctx *gin.Context) {
	var req WebUIListBackgroundRequest
	var resp []WebUIBackgroundEntry

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dictionary := gamedata.DictionaryByLanguage(req.Language)
	session := ctx.MustGet("session").(*userdata.Session)

	for id, background := range gamedata.Instance.CustomBackground {
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

	jsonBytes, err := json.Marshal(resp)
	utils.CheckErr(err)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, string(jsonBytes))
}

func init() {
	server.AddHandler("/webui/user", "GET", "/background", backgroundList)
}
