package emblem

import (
	"encoding/json"

	"elichika/internal/client"
	"elichika/internal/client/request"
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/gin-gonic/gin"
)

func fetchEmblemById(ctx *gin.Context) {
	req := request.EmblemSearchUserIdRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	otherUserSession := userdata.GetSessionWithSharedDb(ctx, req.UserId, session)
	otherUserSession.PopulateUserModelField("UserEmblemByEmblemId")
	resp := response.FetchEmblemResponse{
		UserModel: &session.UserModel,
	}
	for _, emblem := range otherUserSession.UserModel.UserEmblemByEmblemId.Map {
		resp.EmblemIsNewDataList.Append(client.EmblemIsNewData{
			EmblemMasterId: emblem.EmblemMId,
			IsNew:          emblem.IsNew,
		})
	}

	common.JsonResponse(ctx, &resp)
}

func init() {
	server.AddHandler("/", "POST", "/emblem/fetchEmblemById", fetchEmblemById)
}
