package user

import (
	"elichika/internal/server"
	"elichika/internal/userdata"
	"elichika/internal/utils"
	"elichika/internal/webui/response"

	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Unauthenticated endpoint that gives basic information about users for selection in the WebUI login dropdown.
func listUserAccounts(ctx *gin.Context) {
	resp := response.WebUIAccountListResponse{}

	db := userdata.Engine.NewSession()
	err := db.Begin()
	utils.CheckErr(err)

	err = db.Table("u_status").Find(&resp.Accounts)
	utils.CheckErr(err)

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
