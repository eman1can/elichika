package auth

import (
	"net/http"

	"elichika/internal/server"

	"github.com/gin-gonic/gin"
)

// logout handles POST /webui/logout.
func logout(ctx *gin.Context) {
	if err := adapter.Destroy(ctx); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error": nil})
}

func init() {
	server.AddHandler("/webui", "POST", "/logout", logout)
}
