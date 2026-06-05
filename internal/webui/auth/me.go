package auth

import (
	"net/http"

	"elichika/internal/server"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

type UserAccountInfo struct {
	UserId         int32  `json:"user_id"`
	Name           string `json:"name"`
	Nickname       string `json:"nickname"`
	LastLoginAt    int64  `json:"last_login_at"`
	Rank           int32  `json:"rank"`
	Exp            int32  `json:"exp"`
	Message        string `json:"message"`
	ImageAssetPath string `json:"image_asset_path"`
}

// me handles GET /webui/me. Returns the current session state so the frontend
// can detect server restarts or session expiry.
func me(ctx *gin.Context) {
	if isAdmin, _ := adapter.Get(ctx, sessionAdminKey).(bool); isAdmin {
		ctx.JSON(http.StatusOK, gin.H{"authenticated": true, "admin": true, "developer": false})
		return
	}
	if isDeveloper, _ := adapter.Get(ctx, sessionDeveloperKey).(bool); isDeveloper {
		ctx.JSON(http.StatusOK, gin.H{"authenticated": true, "admin": false, "developer": true})
		return
	}
	raw := adapter.Get(ctx, sessionUserIDKey)
	if raw != nil {
		if userId, ok := sessionInt(raw); ok {
			session := userdata.GetSession(ctx, int32(userId))
			if session == nil {
				ctx.JSON(http.StatusOK, gin.H{"authenticated": false, "admin": false, "developer": false})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{"authenticated": true, "admin": false, "developer": false, "user": UserAccountInfo{
				UserId:         int32(userId),
				Name:           session.UserStatus.Name.DotUnderText,
				Nickname:       session.UserStatus.Nickname.DotUnderText,
				LastLoginAt:    session.UserStatus.LastLoginAt,
				Rank:           session.UserStatus.Rank,
				Exp:            session.UserStatus.Exp,
				Message:        session.UserStatus.Message.DotUnderText,
				ImageAssetPath: session.Gamedata.Card[session.UserStatus.RecommendCardMasterId].IdolAppearance.ThumbnailAssetPath,
			}})
			session.Close()
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"authenticated": false})
}

func init() {
	server.AddHandler("/webui", "GET", "/me", me)
}
