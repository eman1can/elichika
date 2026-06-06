package auth

import (
	"net/http"
	"strconv"
	"time"

	"elichika/internal/config"
	"elichika/internal/server"

	ginadapter "github.com/39george/scs_gin_adapter"
	"github.com/alexedwards/scs/v2"
	"github.com/gin-gonic/gin"
)

const (
	sessionUserIDKey    = "user_id"
	sessionAdminKey     = "webui_is_admin"
	sessionDeveloperKey = "webui_is_developer"
)

var adapter *ginadapter.GinAdapter

// RequireLogin aborts with 401 if the request has no valid webui session.
// Sets "webui_user_id" (int) in the Gin context for user sessions, or
// "webui_admin" (bool true) for admin sessions.
func RequireLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if isAdmin, _ := adapter.Get(c, sessionAdminKey).(bool); isAdmin {
			c.Set("webui_admin", true)
			c.Next()
			return
		}
		if isDeveloper, _ := adapter.Get(c, sessionDeveloperKey).(bool); isDeveloper {
			c.Set("webui_developer", true)
			c.Next()
			return
		}
		raw := adapter.Get(c, sessionUserIDKey)
		if raw == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		userId, ok := sessionInt(raw)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("user_id", userId)
		c.Next()
	}
}

func sessionInt(v interface{}) (int, bool) {
	switch n := v.(type) {
	case int:
		return n, true
	case int64:
		return int(n), true
	case string:
		i, err := strconv.Atoi(n)
		return i, err == nil
	}
	return 0, false
}

func init() {
	sm := scs.New()
	sm.Lifetime = 24 * time.Hour
	sm.Cookie.Name = "webui_session"
	sm.Cookie.HttpOnly = true
	sm.Cookie.Persist = true
	sm.Cookie.Domain = *config.Conf.ServerAddress
	sm.Store = initSessionStore()

	adapter = ginadapter.New(sm)

	// Apply LoadAndSave to all /webui groups so session data is available
	// everywhere (including the login endpoint and agnostic endpoints).
	server.AddInitialHandler("/webui", adapter.LoadAndSave)
	server.AddInitialHandler("/webui/user", adapter.LoadAndSave)
	// Gate all user-specific endpoints behind a valid webui session.
	// userInitial (from initial.go) appends after these, so order is preserved.
	server.AddInitialHandler("/webui/user", userInitial)
}
