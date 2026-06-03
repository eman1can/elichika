package user

import (
	"encoding/json"
	"net/http"

	"elichika/internal/client"
	"elichika/internal/gamedata"
	"elichika/internal/server"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

type WebUIUserAccountInfo struct {
	UserId         int32  `json:"user_id"`
	Name           string `json:"name"`
	Nickname       string `json:"nickname"`
	LastLoginAt    int64  `json:"last_login_at"`
	Rank           int32  `json:"rank"`
	Exp            int32  `json:"exp"`
	Message        string `json:"message"`
	ImageAssetPath string `json:"image_asset_path"`
}

type WebUIAccountListRequest struct {
	Language string `json:"l"`
	Search   string `json:"s"`
}

// Unauthenticated endpoint that gives basic information about users for selection in the WebUI login dropdown.
func listUserAccounts(ctx *gin.Context) {
	var req WebUIAccountListRequest
	var resp []WebUIUserAccountInfo

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := userdata.Engine.NewSession()
	err := db.Begin()
	utils.CheckErr(err)

	var status []client.UserStatus
	err = db.Table("u_status").Where("name LIKE ?", "%"+req.Search+"%").Find(&status)
	utils.CheckErr(err)

	for _, user := range status {
		resp = append(resp, WebUIUserAccountInfo{
			UserId:         user.UserId,
			Name:           user.Name.DotUnderText,
			Nickname:       user.Nickname.DotUnderText,
			LastLoginAt:    user.LastLoginAt,
			Rank:           user.Rank,
			Exp:            user.Exp,
			Message:        user.Message.DotUnderText,
			ImageAssetPath: gamedata.Instance.Card[user.RecommendCardMasterId].IdolAppearance.ThumbnailAssetPath,
		})
	}

	err = db.Close()
	utils.CheckErr(err)

	jsonBytes, err := json.Marshal(resp)
	utils.CheckErr(err)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, string(jsonBytes))
}

func init() {
	server.AddHandler("/webui/user", "GET", "/list_accounts", listUserAccounts)
}
