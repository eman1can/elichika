package admin

import (
	"crypto/rand"

	"elichika/internal/config"
	"elichika/internal/server"
	"elichika/internal/utils"
)

var adminSessionKey []byte

func randomKey() []byte {
	// random 32 bytes
	b := make([]byte, 32)
	_, err := rand.Read(b)
	utils.CheckErr(err)
	return b
}

func newSessionKey() {
	adminSessionKey = randomKey()
}

func init() {
	newSessionKey()
	server.AddTemplates(config.RootPath + "internal/webui/admin/logged_in_admin.html")
}
