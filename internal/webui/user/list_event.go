package user

import (
	"fmt"
	"net/http"

	"elichika/internal/enum"
	"elichika/internal/gamedata"
	"elichika/internal/server"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

type WebUIListEventEntry struct {
	EventId        int32  `xorm:"pk 'id'" json:"id"`
	EventType      int32  `xorm:"event_type" json:"event_type"`
	Title          string `json:"title"`
	ImageAssetPath string `json:"image_asset_path"`
	ReleaseOrder   int32  `json:"release_order"`
}

func listEvent(ctx *gin.Context) {
	var resp []WebUIListEventEntry

	session := ctx.MustGet("session").(*userdata.Session)
	dictionary := ctx.MustGet("dictionary").(*gamedata.Dictionary)

	for _, event := range session.Gamedata.Event {
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

	ctx.JSON(http.StatusOK, resp)
}

func init() {
	server.AddHandler("/webui/user", "GET", "/list_event", listEvent)
}
