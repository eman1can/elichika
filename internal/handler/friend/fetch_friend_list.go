package friend

import (
	"elichika/internal/client/response"
	"elichika/internal/enum"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_social"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

func fetchFriendList(ctx *gin.Context) {
	session := ctx.MustGet("session").(*userdata.Session)

	common.JsonResponse(ctx, response.FriendListResponse{
		SuccessType:    enum.FriendSuccessTypeNoProblem,
		FriendViewList: user_social.GetFriendViewList(session),
	})
}

func init() {
	server.AddHandler("/", "POST", "/friend/fetchFriendList", fetchFriendList)
}
