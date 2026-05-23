package present

import (
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_present"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

func fetch(ctx *gin.Context) {
	// there is no request body

	session := ctx.MustGet("session").(*userdata.Session)
	resp := response.FetchPresentResponse{
		PresentItems:        user_present.FetchPresentItems(session),
		PresentHistoryItems: user_present.FetchPresentHistoryItems(session),
	}

	session.Finalize() // this is because the fetch request can cause server to delete expired item
	resp.PresentCount = user_present.FetchPresentCount(session)
	common.JsonResponse(ctx, &resp)
}

func init() {
	server.AddHandler("/", "POST", "/present/fetch", fetch)
}
