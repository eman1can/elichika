package item_detail

import (
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/time"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

// TODO(extra): Implement live campaign
func fetchItemDetailRelateLiveList(ctx *gin.Context) {
	// there is no request body

	session := ctx.MustGet("session").(*userdata.Session)
	common.JsonResponse(ctx, response.FetchItemDetailRelateLiveListResponse{
		WeekdayState: time.GetWeekdayState(session),
	})
}

func init() {
	server.AddHandler("/", "POST", "/itemDetail/fetchItemDetailRelateLiveList", fetchItemDetailRelateLiveList)
}
