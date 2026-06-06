package auth

import (
	"crypto/subtle"
	"net/http"
	"strconv"

	"elichika/internal/config"
	"elichika/internal/enum"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_authentication"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password"`
}

// login handles POST /webui/login.
func login(ctx *gin.Context) {
	req := loginRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Username == "admin" {
		if subtle.ConstantTimeCompare([]byte(*config.Conf.AdminPassword), []byte(req.Password)) != 1 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong username or password!"})
			return
		}
		_ = adapter.RenewToken(ctx)
		adapter.Put(ctx, sessionAdminKey, true)
		adapter.Put(ctx, sessionDeveloperKey, false)
		ctx.JSON(http.StatusOK, gin.H{"error": nil})
		return
	}

	if req.Username == "developer" {
		if subtle.ConstantTimeCompare([]byte(*config.Conf.AdminPassword), []byte(req.Password)) != 1 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong username or password!"})
			return
		}
		_ = adapter.RenewToken(ctx)
		adapter.Put(ctx, sessionDeveloperKey, true)
		adapter.Put(ctx, sessionAdminKey, false)
		ctx.JSON(http.StatusOK, gin.H{"error": nil})
		return
	}

	userID, err := strconv.Atoi(req.Username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID!"})
		return
	}

	session := userdata.GetSession(ctx, int32(userID))
	if session == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User doesn't exist!"})
		return
	}
	defer session.Close()

	if !user_authentication.CheckPassWord(session, req.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password!"})
		return
	}

	if session.UserStatus.TutorialPhase != enum.TutorialPhaseTutorialEnd {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Finish the tutorial (in game) first before using the WebUI!"})
		return
	}

	_ = adapter.RenewToken(ctx)
	adapter.Put(ctx, sessionUserIDKey, userID)
	adapter.Put(ctx, sessionAdminKey, false)
	adapter.Put(ctx, sessionDeveloperKey, false)

	ctx.JSON(http.StatusOK, gin.H{"error": nil})
}

func init() {
	server.AddHandler("/webui", "POST", "/login", login)
}
