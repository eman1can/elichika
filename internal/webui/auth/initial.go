package auth

import (
	"net/http"

	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

func userInitial(ctx *gin.Context) {
	raw := adapter.Get(ctx, sessionUserIDKey)
	if raw == nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	userId, ok := sessionInt(raw)
	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.Set("user_id", userId)

	var session *userdata.Session
	defer func() { session.Close() }()

	session = userdata.GetSession(ctx, int32(userId))

	ctx.Set("session", session)
	ctx.Next()
}
