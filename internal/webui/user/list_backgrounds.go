package user

import (
	"encoding/json"
	"net/http"
	"sort"

	"elichika/internal/gamedata"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_custom_background"
	"elichika/internal/userdata"
	"elichika/internal/utils"
	"elichika/internal/webui/request"
	"elichika/internal/webui/response"

	"github.com/gin-gonic/gin"
)

func backgroundList(ctx *gin.Context) {
	resp := response.WebUIBackgroundListResponse{}
	req := request.WebUILanguageRequest{}
	err := ctx.ShouldBindQuery(&req)
	utils.CheckErr(err)

	dictionary := gamedata.DictionaryByLanguage(req.Language)
	session := ctx.MustGet("session").(*userdata.Session)

	for id, background := range gamedata.Instance.CustomBackground {
		userBackground := user_custom_background.GetUserCustomBackground(session, id)
		entry := response.WebUIBackgroundEntry{
			Id:           id,
			Name:         dictionary.Resolve(background.Name),
			DisplayOrder: background.DisplayOrder,
			Owned:        !userBackground.IsNew,
		}
		resp = append(resp, entry)
	}

	sort.Slice(resp, func(i, j int) bool {
		if resp[i].DisplayOrder != resp[j].DisplayOrder {
			return resp[i].DisplayOrder < resp[j].DisplayOrder
		}
		return resp[i].Id < resp[j].Id
	})

	jsonBytes, err := json.Marshal(resp)
	utils.CheckErr(err)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, string(jsonBytes))
}

func init() {
	server.AddHandler("/webui/user", "GET", "/background", backgroundList)
}
