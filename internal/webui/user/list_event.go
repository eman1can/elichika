package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"elichika/internal/enum"
	"elichika/internal/gamedata"
	"elichika/internal/server"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

type WebUIListEventRequest struct {
	Language string `form:"l" json:"l"`
}

type WebUIListEventEntry struct {
	EventId        int32  `xorm:"pk 'id'" json:"id"`
	EventType      int32  `xorm:"event_type" json:"event_type"`
	Title          string `json:"title"`
	ImageAssetPath string `json:"image_asset_path"`
	ReleaseOrder   int32  `json:"release_order"`
}

func listEvent(ctx *gin.Context) {
	var req WebUIListEventRequest
	var resp []WebUIListEventEntry

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dictionary := gamedata.DictionaryByLanguage(req.Language)

	for _, event := range gamedata.Instance.Event {
		var name string
		if event.EventType == enum.EventTypeMining {
			name = fmt.Sprintf("m.event_mining_title_%d", event.EventId)
		} else {
			name = fmt.Sprintf("m.event_marathon_title_%d", event.EventId)
		}

		resp = append(resp, WebUIListEventEntry{
			EventId:        event.EventId,
			EventType:      event.EventType,
			Title:          dictionary.Resolve(name),
			ImageAssetPath: *event.BannerNoticeLargeAssetPath,
			ReleaseOrder:   event.ReleaseOrder,
		})
	}

	jsonBytes, err := json.Marshal(resp)
	utils.CheckErr(err)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, string(jsonBytes))
}

func init() {
	server.AddHandler("/webui/user", "GET", "/list_event", listEvent)
}
