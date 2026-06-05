package webui

import (
	"net/http"

	"elichika/internal/client"
	"elichika/internal/server"
	"elichika/internal/userdata"
	"elichika/internal/utils"
	"elichika/internal/webui/auth"

	"github.com/gin-gonic/gin"
)

type AccountListRequest struct {
	Search string `json:"s"`
}

// Unauthenticated endpoint that gives basic information about users for selection in the WebUI login dropdown.
func listUserAccounts(ctx *gin.Context) {
	var req AccountListRequest
	var resp []auth.UserAccountInfo

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := userdata.GetPlainSession(ctx)

	var status []client.UserStatus
	err := session.Db.Table("u_status").Where("name LIKE ?", "%"+req.Search+"%").Find(&status)
	utils.CheckErr(err)

	for _, user := range status {
		resp = append(resp, auth.UserAccountInfo{
			UserId:         user.UserId,
			Name:           user.Name.DotUnderText,
			Nickname:       user.Nickname.DotUnderText,
			LastLoginAt:    user.LastLoginAt,
			Rank:           user.Rank,
			Exp:            user.Exp,
			Message:        user.Message.DotUnderText,
			ImageAssetPath: session.Gamedata.Card[user.RecommendCardMasterId].IdolAppearance.ThumbnailAssetPath,
		})
	}

	session.Close()

	ctx.JSON(http.StatusOK, resp)
}

func init() {
	server.AddHandler("/webui", "GET", "/accounts", listUserAccounts)
}
