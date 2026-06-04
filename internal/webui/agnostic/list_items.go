package agnostic

import (
	"encoding/json"
	"net/http"

	"elichika/internal/gamedata"
	"elichika/internal/server"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

type WebUIItem struct {
	ContentType int32  `json:"content_type"`
	ContentId   int32  `json:"content_id"`
	Name        string `json:"name"`
	AssetPath   string `json:"asset_path"`
}

type WebUIItemListRequest struct {
	ContentType int32  `form:"type"`
	Language    string `form:"l"`
}

func listItems(ctx *gin.Context) {
	var req WebUIItemListRequest
	var resp []WebUIItem

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dictionary := gamedata.DictionaryByLanguage(req.Language)

	for itemId, content := range gamedata.Instance.Content[req.ContentType] {
		resp = append(resp, WebUIItem{
			ContentType: req.ContentType,
			ContentId:   itemId,
			Name:        dictionary.Resolve(content.Name),
			AssetPath:   content.AssetPath,
		})
	}

	jsonBytes, err := json.Marshal(resp)
	utils.CheckErr(err)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, string(jsonBytes))
}

func init() {
	server.AddHandler("/webui", "GET", "/item", listItems)
}
